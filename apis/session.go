package apis

import (
	"github.com/gin-gonic/gin"
)

func NewSession(c *gin.Context) {
	// var user model.User
	// err := c.ShouldBindJSON(&user)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code": 2001,
	// 		"msg":  "无效的参数",
	// 	})
	// 	return
	// }

	// if users, err := user.Users(); err != nil || len(users) == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"code": 2003,
	// 		"msg":  "用户不存在",
	// 	})
	// 	return
	// }

	// token, _ := middleware.GenToken(user.ID)
	// c.JSON(http.StatusOK, gin.H{
	// 	"code":  2000,
	// 	"msg":   "success",
	// 	"id":    user.ID,
	// 	"token": token,
	// })
}
