package userservice

import "banana/serializer"

type UserUpdateService struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u *UserUpdateService) Update() *serializer.Response {
	return &serializer.Response{
		Code: 200,
		Msg:  "更新成功",
		Data: "",
	}

}
