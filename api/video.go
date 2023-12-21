package api

import (
	"github.com/gin-gonic/gin"

	"github.com/hd2yao/giligili/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
