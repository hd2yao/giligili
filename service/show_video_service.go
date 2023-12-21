package service

import (
    "github.com/hd2yao/giligili/model"
    "github.com/hd2yao/giligili/serializer"
)

type ShowVideoService struct{}

func (service *ShowVideoService) Show(id string) serializer.Response {
    var video model.Video
    if err := model.DB.First(&video, id).Error; err != nil {
        return serializer.Response{
            Code:  404,
            Msg:   "视频不存在",
            Error: err.Error(),
        }
    }

    return serializer.Response{
        Code: 200,
        Msg:  "视频详情查询成功",
        Data: serializer.BuildVideo(video),
    }
}
