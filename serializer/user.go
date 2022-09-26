package serializer

import (
	"banana/model"
	"time"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedOn string `json:"created_on"`
}

// 序列化用户
func BuildUser(user *model.User) User {
	str_time := time.Unix(user.CreatedOn, 0).Format("2006-01-02 15:04")
	return User{
		ID:        uint(user.ID),
		UserName:  user.Username,
		NickName:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedOn: str_time,
	}
}
