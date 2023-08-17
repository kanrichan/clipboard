package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/entity"
	"github.com/kanrichan/clipboard/model"
	"github.com/kanrichan/clipboard/router/middleware"
)

func NewSession(c *gin.Context) {
	var request entity.NewSessionRequest
	var response entity.NewSessionResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	id, err := model.VerifyUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
		return
	}
	if id == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidPasswordCode, "message": ErrInvalidPassword})
		return
	}
	token, err := middleware.GenToken(id, request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
		return
	}
	response.ID = id
	response.Username = request.Username
	response.Token = token
	c.JSON(200, gin.H{"code": 0, "message": "登录成功", "data": response})
}

func GetSession(c *gin.Context) {
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
	c.JSON(200, gin.H{"code": 0, "message": ""})
}

func PutSession(c *gin.Context) {
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
	c.JSON(200, gin.H{"code": 0, "message": ""})
}

func DelSession(c *gin.Context) {
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
	c.JSON(200, gin.H{"code": 0, "message": "注销成功"})
}
