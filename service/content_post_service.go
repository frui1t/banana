package service

import (
	"banana/model"
	"banana/serializer"
	"context"
)

type ContentPostService struct {
	Content string `json:"content" from:"content" binding:"required"`
}

func (c *ContentPostService) ContentPost(ctx context.Context, uid uint) *serializer.Response {
	con := model.Content{
		User_id: uid,
		Content: c.Content,
	}
	if err := model.DB.Model(&model.Content{}).Create(&con).Error; err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "发表动态失败",
		}
	}
	return &serializer.Response{
		Code:  200,
		Msg:   "success",
		Data:  "发表动态成功",
		Error: "",
	}
}
