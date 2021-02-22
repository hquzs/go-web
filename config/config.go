package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"hquzs/go-web/util"

	"gopkg.in/yaml.v2"
)

const (
	// DefaultConnections 128
	DefaultConnections = 128
)

var (
	log = util.NewZeroLog("web", "config")
)

// Config : Used for config txserver's zone, pk, issuing pk
type Config struct {
	LogLevel        string     `yaml:"LogLevel"`        // log level, Debug means debug while others mean release
	ConnectionLimit int        `yaml:"ConnectionLimit"` // client connect limit, more than 256
	HTTPConfig      HTTPConfig `yaml:"HTTPConfig"`      // txserver api config
	TLS             TLS        `yaml:"TLS"`             // tls config

	// Database        Database   `yaml:"Database"`        // database rpc config
	// Consensus        string           `yaml:"Consensus"`
	// Version              int32            // currency version
}

// HTTPConfig api server config
type HTTPConfig struct {
	Port int    `yaml:"Port"` // API serve on port
	Host string `yaml:"Host"` // API serve host
}

// TLS tls config
type TLS struct {
	Enable      bool             `yaml:"Enable"`
	CA          string           `yaml:"CA"`
	CAPool      *x509.CertPool   `yaml:"-"`
	Cert        string           `yaml:"Cert"`
	Certificate *tls.Certificate `yaml:"-"`
	Key         string           `yaml:"Key"`
}

// LoadConfig loads config from config file
func LoadConfig(cfgFile string) (*Config, error) {
	log.Info("Load cfgFile from: ", cfgFile)
	if cfgFile == "" {
		return nil, errors.New("No config file setting")
	}
	cfgBytes, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(cfgBytes, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.ConnectionLimit == 0 {
		cfg.ConnectionLimit = DefaultConnections
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	return cfg, nil

}
