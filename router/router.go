package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/apis"
)

func Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/user", apis.NewUser)
	api.POST("/session", apis.NewSession)

	auth := api.Group("")
	auth.DELETE("/user/:id", apis.DelUser)
	auth.GET("/user", apis.GetAllUsers)

	return router
}
