package userservice

import (
	"banana/model"
	"banana/serializer"
	"context"
	"strconv"
)

type UserLogoutService struct{}

func (u *UserLogoutService) Logout(ctx context.Context, uid uint) *serializer.Response {
	_, err := model.RDB.Del(context.Background(), strconv.FormatUint(uint64(uid), 10)).Result()
	if err != nil {
		return &serializer.Response{
			Code:  200,
			Msg:   "err",
			Data:  "redis删除失败",
			Error: "",
		}
	}
	return &serializer.Response{
		Code:  200,
		Msg:   "success",
		Data:  "退出登录成功",
		Error: "",
	}
}
