package api

import (
	"github.com/gin-gonic/gin"
)
func (s *service)TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		actId := c.GetHeader("x-actid")
		token := c.GetHeader("x-token")

		if actId != "123" || token != "456"{
			c.JSON(-111, map[string]interface{}{"msg":"校验失败"})
			c.Abort()
			return
		}
		c.Next()
	}
}