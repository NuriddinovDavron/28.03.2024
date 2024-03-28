package storage

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Product struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

type Message struct {
	Message string `json:"message"`
}