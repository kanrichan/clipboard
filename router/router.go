package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/apis"
)

func Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/user/new", apis.NewUser)
	// api.DELETE("/sessions", apis.Store)

	return router
}
