package service

import (
	"banana/model"
	"banana/serializer"
)

type UserPostService struct {
	Username string `json:"username" from:"username" binding:"required"`
}

/*
func (u *UserPostService) UserPost(ctx context.Context, uid uint, id uint) *serializer.Response {
	err := model.DB.Model(&model.User{}).Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return &serializer.Response{
			Code:  400,
			Data:  "post err",
			Msg:   "post err",
			Error: "post err",
		}
	}
	return &serializer.Response{
		Code:  200,
		Data:  serializer.BuildUser(&user),
		Msg:   "success",
		Error: "err",
	}
}
*/

func (u *UserPostService) UserPost() *serializer.Response {
	err := model.DB.Model(&model.User{}).Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return &serializer.Response{
			Code:  400,
			Data:  "post err",
			Msg:   "post err",
			Error: "post err",
		}
	}
	return &serializer.Response{
		Code:  200,
		Data:  serializer.BuildUser(&user),
		Msg:   "success",
		Error: "err",
	}
}
