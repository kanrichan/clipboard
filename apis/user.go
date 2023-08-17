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
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	id, err := model.NewUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidUsernameCode, "message": ErrInvalidUsername})
		return
	}
	response.ID = id
	c.JSON(200, gin.H{"code": 0, "message": "注册成功", "data": response})
}

func GetUserByID(c *gin.Context) {
	var response entity.GetUserByIDResponse
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	if cid, ok := c.Get("id"); !ok || cid != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": ErrUnAuthorizedCode, "message": ErrUnAuthorized})
		return
	}
	user, err := model.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	if cid, ok := c.Get("id"); !ok || cid != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": ErrUnAuthorizedCode, "message": ErrUnAuthorized})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	if err := model.UpdUserByID(id, request.Username, request.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidUserIDCode, "message": ErrInvalidUserID})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "更新成功"})
}

func DelUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	if cid, ok := c.Get("id"); !ok || cid != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": ErrUnAuthorizedCode, "message": ErrUnAuthorized})
		return
	}
	if err := model.DelUserByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidUserIDCode, "message": ErrInvalidUserID})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "销户成功"})
}
