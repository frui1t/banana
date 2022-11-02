package userservice

import (
	"banana/model"
	"banana/serializer"
)

type UserUpdate struct {
	Nickname string `json:"nickname"`
	Password string `json:"passwod"`
	Avatar   string `json:"avatar"`
}

func (u *UserUpdate) Update(uid uint) *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", u.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40000,
			Msg:  "昵称已被占用",
		}
	}

	return &serializer.Response{
		Code: 200,
		Msg:  "更新成功",
	}
}
