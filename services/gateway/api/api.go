package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func (s *service) Test(c *gin.Context) {
	fmt.Println("test!!")
}