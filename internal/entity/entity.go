package entity

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Item struct {
	Id int `json:"id" validate:"required,min=1" db:"id"`
	Title string `json:"title" validate:"required" db:"title"`
	Description string `json:"description" db:"description"`
}
