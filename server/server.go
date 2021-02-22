package server

import (
	"context"
	"fmt"
	"hquzs/go-web/config"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/netutil"
)

// WebVersion ...
const WebVersion = "v1"

// WebServer web-server http api service
type WebServer struct {
	Engine         *gin.Engine
	Config         config.HTTPConfig
	srv            *http.Server
	MaxConnections int
	ln             net.Listener
}

//StartServer start server with cfg
func StartServer(cfgFile string) error {
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		return fmt.Errorf("Load config failed: %v", err)
	}
	s, err := NewWebServer(cfg)
	if err != nil {
		return fmt.Errorf("NewWebServer failed: %v", err)
	}
	err = s.Start()
	if err != nil {
		return fmt.Errorf("Start WebServer failed: %v", err)
	}
	return nil
}

// NewWebServer init WebServer with config
func NewWebServer(cfg *config.Config) (*WebServer, error) {
	log.Info("init new WebServer")

	if strings.ToLower(cfg.LogLevel) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	setLevel(cfg.LogLevel)

	webServer := &WebServer{
		Engine:         gin.New(),
		Config:         cfg.HTTPConfig,
		MaxConnections: cfg.ConnectionLimit,
	}
	webServer.Engine.Use(gin.Recovery())

	webServer.srv = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTPConfig.Host, cfg.HTTPConfig.Port),
		Handler:      webServer.Engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := webServer.initRouter()
	if err != nil {
		log.Error("Init router failed: ", err)
		return nil, err
	}
	return webServer, nil
}

func (s *WebServer) initRouter() error {
	v := s.Engine.Group(WebVersion)

	{
		v.GET("hello", s.hello)
	}
	return nil
}

// Start start web-server http api
func (s *WebServer) Start() error {
	log.Info("Call start WebServer,listen：", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	var err error
	var ln net.Listener

	ln, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		log.Error("HTTP listen failed")
		return err
	}
	defer ln.Close()

	limit := s.MaxConnections
	if limit == 0 {
		limit = config.DefaultConnections
	}
	if limit > 0 {
		log.Infof("Gin server limitlistener %d", limit)
		ln = netutil.LimitListener(ln, limit)
	}

	go func() {
		err := s.srv.Serve(ln)
		if err != nil && err != http.ErrServerClosed {
			log.Error("Http Serve failed: ", err)
		}
	}()

	log.Info("Server start succeed.")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	}
	log.Info("Server exiting")

	return nil
}
