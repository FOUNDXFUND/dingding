package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	ctx *gin.Context
}

func NewResp(ctx *gin.Context) *Resp {
	return &Resp{ctx}
}

func (r *Resp) 	SuccWithData(data interface{}) {
	r.ctx.JSON(http.StatusOK, map[string]interface{}{"code":0, "data":data})
	return
}

func (r *Resp) SuccEmpty() {
	r.ctx.JSON(http.StatusOK, map[string]interface{}{"code":0})
	return
}


func CommonErr(code int,data interface{}) (comm map[string]interface{}){
	comm = map[string]interface{}{
		"code":code,
		"data":data,
	}
	return
}

func (r *Resp) Error(code int,data interface{}) {
	j := CommonErr(code,data)
	r.ctx.JSON(http.StatusOK, j)
	return
}