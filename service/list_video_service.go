package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video

	err := model.DB.Find(&videos).
		Limit(service.Limit).Offset(service.Start).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
