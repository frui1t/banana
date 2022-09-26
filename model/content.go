package model

type Content struct {
	*Model
	User_id uint   `json:"user_id"`
	Content string `json:"content"`
}
