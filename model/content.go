package model

type Content struct {
	*Model
	User_id string `json:"user_id"`
	Content string `json:"content"`
}
