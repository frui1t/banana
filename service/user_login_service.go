package service

import (
	"banana/model"
	"banana/serializer"
	"banana/util"
)

type UserLoginService struct {
	Username string `json:"username" from:"username" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}

var user model.User

func (u *UserLoginService) valid() *serializer.Response {

	if err := model.DB.Model(&model.User{}).Where("username = ?", u.Username).Find(&user).Error; err != nil {
		return &serializer.Response{
			Code:  40000,
			Data:  "账号或者密码错误",
			Msg:   "login err",
			Error: "err",
		}
	}
	if !user.CheckPassword(u.Password) {
		return &serializer.Response{
			Code:  40000,
			Data:  "账号或者密码错误",
			Msg:   "login err",
			Error: "err",
		}
	}

	return nil
}

func (u *UserLoginService) Login() *serializer.Response {
	if err := u.valid(); err != nil {
		return err
	}
	//token签发
	accesstoken, err := util.GenerateAccessToken(uint(user.ID), u.Username, 0)
	if err != nil {
		return &serializer.Response{
			Code: 200,
			Data: "",
			Msg:  "accesstoken 签发失败",
		}
	}

	refreshtoken, err := util.GenerateRefreshToken(uint(user.ID), u.Username, 0)
	if err != nil {
		return &serializer.Response{
			Code: 200,
			Data: "",
			Msg:  "refreshtoken 签发失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Data: serializer.TokenData{User: serializer.BuildUser(&user), Access_token: accesstoken, Refresh_token: refreshtoken},
		Msg:  "login success",
	}
}
