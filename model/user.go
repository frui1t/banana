package model

import (
	"banana/util"

	"github.com/sirupsen/logrus"
)

type User struct {
	*Model
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u *User) SetPassword(password string) error {
	newPassword, err := util.EncodeMD5(password)
	if err != nil {
		logrus.Error(err)
		return err
	}
	u.Password = newPassword
	return nil
}

func (u *User) CheckPassword(password string) bool {
	checkPassword, err := util.EncodeMD5(password)
	if err != nil {
		logrus.Error(err)
		return false
	}
	if checkPassword != u.Password {
		return false
	}
	return true

}
