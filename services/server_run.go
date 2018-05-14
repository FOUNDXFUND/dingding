package services

import (
	"runtime"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"dingding/conf"
	"dingding/utils"
	"dingding/common/errors"
)

type Service interface {
	Name() string
	Version() string
	Config(conf *conf.ServiceConf, r *gin.Engine) error
	Start()
}

func RunService(service Service) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("service:", service.Name())
	conf, err := conf.LoadServiceConf(service.Name())
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(catchError())
	extraProgName := fmt.Sprintf("-version:%s -port:%d", service.Version(), conf.Server.Port)
	utils.SetVersion(extraProgName)
	service.Config(conf,r)
	service.Start()
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	r.Run(addr)
}

func catchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *errors.Error:
					e := err.(*errors.Error)
					c.JSON(e.HTTPStatus, e)
					c.Abort()
				default:
					//c.JSON(http.StatusInternalServerError, common.CommonErr(-500, fmt.Sprintf("系统错误:%v",err)))
					panic(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}