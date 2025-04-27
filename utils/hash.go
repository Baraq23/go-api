package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+"Wolfgang"), 14)
	return string(bytes), err
}