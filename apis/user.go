package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/entity"
	"github.com/kanrichan/clipboard/model"
)

func NewUser(c *gin.Context) {
	var request entity.NewUserRequest
	var response entity.NewUserResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	id, err := model.NewUserByUsername(request.Username, request.Password)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "用户名已存在"})
		return
	}
	response.ID = id
	c.JSON(200, gin.H{"code": 0, "message": "注册成功", "data": response})
}

func DelUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
		return
	}
	if err := model.DelUserByID(id); err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "用户ID不存在"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "销户成功"})
}

func GetAllUsers(c *gin.Context) {
	var response entity.GetAllUsersResponse
	users, err := model.GetAllUsers()
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "message": "获取用户列表出错"})
		return
	}
	for i := range users {
		response.Users = append(response.Users, struct {
			ID       int64  "json:\"id\""
			Username string "json:\"username\""
		}{ID: users[i].ID, Username: users[i].Username})
	}
	c.JSON(200, gin.H{"code": 0, "message": "获取成功", "data": response})
}
