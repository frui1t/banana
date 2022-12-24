package userservice

import (
	"banana/model"
	"banana/serializer"
	"banana/util"

	"github.com/sirupsen/logrus"
)

type UserUpdateService struct {
	Nickname string `json:"nickname" from:"nickname" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}

func (u *UserUpdateService) Update(uid uint) *serializer.Response {

	count := int64(0)
	logrus.Warningln("ttt", u.Nickname)
	model.DB.Model(&model.User{}).Where("nickname = ?", u.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40000,
			Msg:  "昵称已被占用",
		}
	}
	if len(u.Password) != 0 {
		var err error
		u.Password, err = util.EncodeMD5(u.Password)
		if err != nil {
			return &serializer.Response{
				Code: 200,
				Msg:  "更新密码失败",
				Data: "",
			}
		}
	}

	err := model.DB.Table("user").Where("id = ?", uid).Updates(u).Error
	if err != nil {
		return &serializer.Response{
			Code: 200,
			Msg:  "更新失败",
			Data: "",
		}

	}
	return &serializer.Response{
		Code: 200,
		Msg:  "更新成功",
		Data: "",
	}

}
