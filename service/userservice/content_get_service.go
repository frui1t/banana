package userservice

import (
	"banana/model"
	"banana/serializer"
)

type ContentGetService struct{}

func (c *ContentGetService) ContentGet() *serializer.Response {
	var contents []*model.Content
	model.DB.Table("content").Find(&contents)
	return &serializer.Response{
		Code:  200,
		Data:  contents,
		Msg:   "",
		Error: "",
	}
}
