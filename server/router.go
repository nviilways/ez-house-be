package server

import (
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {

}

func NewRoute(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
}