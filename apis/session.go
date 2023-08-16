package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/entity"
	"github.com/kanrichan/clipboard/model"
	"github.com/kanrichan/clipboard/router/middleware"
)

func NewSession(c *gin.Context) {
	var request entity.NewSessionRequest
	var response entity.NewSessionResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	id, err := model.VerifyUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "系统错误"})
		return
	}
	if id == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "用户名或密码错误"})
		return
	}
	token, err := middleware.GenToken(id, request.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "系统错误"})
		return
	}
	response.ID = id
	response.Username = request.Username
	response.Token = token
	c.JSON(200, gin.H{"code": 0, "message": "注册成功", "data": response})
}

func DelSession(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "message": "注销成功"})
}
