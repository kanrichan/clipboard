package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/apis"
	"github.com/kanrichan/clipboard/router/middleware"
)

func Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	auth := api.Group("", middleware.AuthSessionMiddle())

	api.POST("/user", apis.NewUser)
	auth.GET("/user/:id", apis.GetUserByID)
	auth.PUT("/user/:id", apis.UpdUserByID)
	auth.DELETE("/user/:id", apis.DelUserByID)

	api.POST("/session", apis.NewSession)
	auth.GET("/session")
	auth.PUT("/session")
	auth.DELETE("/session", apis.DelSession)

	return router
}
