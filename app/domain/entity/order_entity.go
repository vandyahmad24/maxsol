package entity

import (
	"github.com/go-playground/validator/v10"
	formater "github.com/vandyahmad24/validator-formater"
	"time"
)

type OrderInput struct {
	CakeId int `json:"cake_id" validate:"required"`
	Qty    int `json:"qty" validate:"required"`
}

func ValidateInputOrder(input OrderInput) interface{} {
	var errors interface{}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors = formater.FormatErrorValidation(err, "You must complete input")
	}
	return errors
}

type OrderResponseWithoutCake struct {
	Id        int       `json:"id"`
	CakeId    int       `json:"cake_id" `
	Qty       int       `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
}

type OrderInputBulk struct {
	Data []OrderInput `json:"data" validate:"required"`
}

func ValidateInputBulk(input OrderInputBulk) interface{} {
	var errors interface{}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors = formater.FormatErrorValidation(err, "You must complete input")
	}
	return errors
}
