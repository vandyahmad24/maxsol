package entity

import (
	"github.com/go-playground/validator/v10"
	formater "github.com/vandyahmad24/validator-formater"
	"time"
)

type UserInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ValidateUserInput(input UserInput) interface{} {
	var errors interface{}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors = formater.FormatErrorValidation(err, "You must complete input")
	}
	return errors
}

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
