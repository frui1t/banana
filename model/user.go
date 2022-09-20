package model

type User struct {
	*Model
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}
