package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

type CreateAccountRequest struct {
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// function to generate new account
func NewAccount(fistname, lastname, password, email, username string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	return &Account{
		FirstName:         fistname,
		LastName:          lastname,
		Username:          username,
		Email:             email,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now(),
	}, nil
}
