package userservice

import (
	"banana/model"
	"banana/serializer"
	"banana/util"
	"context"
	"strconv"
	"time"
)

type UserLoginService struct {
	Username string `json:"username" from:"username" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}

func (u *UserLoginService) Login() *serializer.Response {
	var user model.User
	if err := model.DB.Table("user").Where("username = ?", u.Username).Take(&user).Error; err != nil {
		return &serializer.Response{
			Code:  40000,
			Data:  "账号错误",
			Msg:   "login err",
			Error: "err",
		}
	}
	if !user.CheckPassword(u.Password) {
		return &serializer.Response{
			Code:  40000,
			Data:  "密码错误",
			Msg:   "login err",
			Error: "err",
		}
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
	refreshtoken, err := util.GenerateRefreshToken(0)
	if err != nil {
		return &serializer.Response{
			Code: 200,
			Data: "",
			Msg:  "refreshtoken 签发失败",
		}
	}
	//logrus.Warnln(refreshtoken)

	err = model.RDB.Set(context.Background(), strconv.FormatInt(user.ID, 10), refreshtoken, time.Hour*24).Err()
	if err != nil {
		return &serializer.Response{
			Code: 200,
			Data: err,
			Msg:  "refreshtoken re dis 签发失败",
		}
	}

	return &serializer.Response{
		Code: 200,
		Data: serializer.TokenData{User: serializer.BuildUser(&user), Access_token: accesstoken},
		Msg:  "login success",
	}
}
