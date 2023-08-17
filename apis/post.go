package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kanrichan/clipboard/entity"
	"github.com/kanrichan/clipboard/model"
)

func NewPost(c *gin.Context) {
	var request entity.NewPostRequest
	var response entity.NewPostResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	cid, ok := c.Get("id")
	if !ok || cid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	var userID = cid.(int64)
	id, err := model.NewPost(userID, request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
		return
	}
	response.ID = id
	c.JSON(200, gin.H{"code": 0, "message": "发表成功", "data": response})
}

func GetPostByID(c *gin.Context) {
	var response entity.GetPostByIDResponse
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	if cid, ok := c.Get("id"); !ok || cid == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": ErrUnAuthorizedCode, "message": ErrUnAuthorized})
		return
	}
	post, err := model.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidPostIDCode, "message": ErrInvalidPostID})
		return
	}
	user, err := model.GetUserByID(post.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
		return
	}
	response.ID = post.ID
	response.Content = post.Content
	response.Author = struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	}{user.ID, user.Username}
	c.JSON(200, gin.H{"code": 0, "message": "获取成功", "data": response})
}

func UpdPostByID(c *gin.Context) {
	var request entity.UpdPostByIDRequest
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	post, err := model.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
		return
	}
	if cid, ok := c.Get("id"); !ok || cid != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": ErrUnAuthorizedCode, "message": ErrUnAuthorized})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	if err := model.UpdPostByID(id, request.Content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidPostIDCode, "message": ErrInvalidPostID})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "更新成功"})
}

func DelPostByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrMissingParameterCode, "message": ErrMissingParameter})
		return
	}
	post, err := model.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": ErrInternalErrorCode, "message": ErrInternalError})
		return
	}
	if cid, ok := c.Get("id"); !ok || cid != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": ErrUnAuthorizedCode, "message": ErrUnAuthorized})
		return
	}
	if err := model.DelPostByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": ErrInvalidPostIDCode, "message": ErrInvalidPostID})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "删除成功"})
}
