package api

import (
	"github.com/gin-gonic/gin"
	"dingding/conf"
	"dingding/services"
	"dingding/common"
)

type service struct {
	conf *conf.ServiceConf
	resp *common.Resp
}

const (
	VERSION = "0.0.1"
)

func NewService() services.Service {
	return new(service)
}

func (s *service) Name() string {
	return "conf_template"
}

func (s *service) Version() string {
	return VERSION
}

func (s *service) Config(conf *conf.ServiceConf, rr *gin.Engine) error {
	s.conf = conf
	rr.Use(func(c *gin.Context) {s.resp = common.NewResp(c)})

	// no auth :
	rr.POST("/login", s.Login)
	rr.Use(s.TokenAuthMiddleware())

	// auth:
	rr.POST("/test", s.Test)
	return nil
}

func (s *service) Start() {

}




