package models

import (
	"errors"
	"goapi/db"
	"goapi/utils"
)

type User struct {
	ID       int64
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = userId

	return nil
}

func (u User) ValidateCridentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	// email not found
	if err != nil {
		return errors.New("Invalid credentials")
	}

	// email found code continues; check if password bound to that email is valid
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
