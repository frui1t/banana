package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) (string, error) {
	m := md5.New()
	_, err := m.Write([]byte(value))
	newPasswd := hex.EncodeToString(m.Sum(nil))
	return newPasswd, err
}
