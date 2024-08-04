package models

import (
	"database/sql"
	"errors"
	"restapi_example/db"
	"restapi_example/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password) 
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
    SELECT id, password FROM users WHERE email =?`

	stmt := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := stmt.Scan(&u.ID, &retrievedPassword)

	if err == sql.ErrNoRows {
		return err
	} else if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("password is not valid")
	}
	return nil
}
