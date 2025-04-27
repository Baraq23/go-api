package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+"Wolfgang"), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+"Wolfgang"))
	return err == nil
}