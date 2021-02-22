package server

import (
	"hquzs/go-web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newResp(err error) gin.H {
	if err != nil {
		var e *util.Error
		var ok bool
		e, ok = err.(*util.Error)
		if !ok {
			e = util.NewError(err, util.UnknownError)
		}
		return gin.H{
			"msg":  e.Err.Error(),
			"code": e.Code,
		}
	}
	return gin.H{
		"code": util.SUCCESS,
	}
}

func bindErrResp(tag string, err *util.Error) gin.H {
	log.Infof("%s Bind request failed: %v", tag, err)
	return gin.H{
		"msg":  "参数解析失败",
		"code": err.Code,
	}
}

func reqErrResp(err string) gin.H {
	return gin.H{
		"msg":  err,
		"code": util.InvalidRequest,
	}
}

func (s *WebServer) hello(c *gin.Context) {
	log.Info("start hello server")
	var err error
	resp := newResp(err)
	resp["msg"] = "hello world"
	c.JSON(http.StatusOK, resp)
	return
}
