package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/apis"
	"github.com/kanrichan/clipboard/router/middleware"
)

func Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	auth := api.Group("", middleware.AuthSession())

	api.POST("/user", apis.NewUser)
	auth.GET("/user/:id", apis.GetUserByID)
	auth.PUT("/user/:id", apis.UpdUserByID)
	auth.DELETE("/user/:id", apis.DelUserByID)

	api.POST("/session", apis.NewSession)
	auth.GET("/session/:id", apis.GetSession)
	auth.PUT("/session/:id", apis.UpdSession)
	auth.DELETE("/session/:id", apis.DelSession)

	auth.POST("/post", apis.NewPost)
	auth.GET("/post/:id", apis.GetPostByID)
	auth.PUT("/post/:id", apis.UpdPostByID)
	auth.DELETE("/post/:id", apis.DelPostByID)

	return router
}
