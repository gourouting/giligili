package api

import (
	"giligili/serializer"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "成功",
	})
}
