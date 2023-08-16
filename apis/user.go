package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/entity"
	"github.com/kanrichan/clipboard/model"
)

func NewUser(c *gin.Context) {
	var request entity.NewUserRequest
	var response entity.NewUserResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	id, err := model.NewUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "用户名已存在"})
		return
	}
	response.ID = id
	c.JSON(200, gin.H{"code": 0, "message": "注册成功", "data": response})
}

func DelUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	if err := model.DelUserByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "用户ID不存在"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "销户成功"})
}

func GetUserByID(c *gin.Context) {
	var response entity.GetUserByIDResponse
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	user, err := model.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "获取用户列表出错"})
		return
	}
	response.ID = user.ID
	response.Username = user.Username
	c.JSON(200, gin.H{"code": 0, "message": "获取成功", "data": response})
}

func UpdUserByID(c *gin.Context) {
	var request entity.UpdUserByIDRequest
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	if err := model.UpdUserByID(id, request.Username, request.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "更新用户信息失败"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "更新成功"})
}
