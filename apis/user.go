package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/model"
)

func NewUser(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	_, err := model.NewUser(request.Username, request.Password)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "用户名已存在"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "注册成功"})
}
