package service

import (
	"github.com/hd2yao/giligili/model"
	"github.com/hd2yao/giligili/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
}

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "视频详情查询成功",
		Data: serializer.BuildVideos(videos),
	}
}
