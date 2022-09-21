package service

import (
	"banana/model"
	"banana/serializer"
)

// 用户注册服务
type UserRegisterService struct {
	Nickname        string `json:"nickname" from:"nickname" binding:"required"`
	Username        string `json:"username" from:"username" binding:"required"`
	Password        string `json:"password" from:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" from:"password_confirm" binding:"required"`
}

func (u *UserRegisterService) valid() *serializer.Response {
	if u.Password != u.PasswordConfirm {
		return &serializer.Response{
			Code: 40000,
			Msg:  "两次输入的密码不相同",
		}
	}
	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", u.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40000,
			Msg:  "昵称已被占用",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40000,
			Msg:  "用户名已被占用",
		}
	}
	return nil
}

func (u *UserRegisterService) Register() *serializer.Response {
	user := model.User{
		Nickname: u.Nickname,
		Username: u.Username,
		Avatar:   "000",
	}
	if err := u.valid(); err != nil {
		return err
	}

	user.SetPassword(u.Password)

	if err := model.DB.Model(&model.User{}).Create(&user).Error; err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "注册失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Msg:  "注册成功",
	}
}
