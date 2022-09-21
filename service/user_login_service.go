package service

import (
	"banana/model"
	"banana/serializer"
)

type UserLoginService struct {
	Username string `json:"username" from:"username" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}

func (u *UserLoginService) valid() *serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("username = ?", u.Username).First(&user).Error; err != nil {
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
	return &serializer.Response{
		Code: 200,
		Data: "登录成功",
		Msg:  "login success",
	}
}
