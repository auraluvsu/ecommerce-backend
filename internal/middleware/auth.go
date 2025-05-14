package middleware

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	Name string
	Pswd string
	Id   string
}

func (u *User) CreateID() (string, error) {
	randBytes := make([]byte, 32)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(randBytes)
	return hex.EncodeToString(hash[:]), nil
}
