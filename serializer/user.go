package serializer

import "banana/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// 序列化用户
func BuildUser(user *model.User) User {
	return User{
		ID:       uint(user.ID),
		UserName: user.Username,
		NickName: user.Nickname,
		Avatar:   user.Avatar,
	}
}
