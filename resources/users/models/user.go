package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             int    `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	HashedPassword string `json:"-"`
}

func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
	return string(hash)
}

func ComparePasswords(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}
