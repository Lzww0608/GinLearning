package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	type User struct {
		ID     int64  `json:"id" validate:"gt=0"`
		Name   string `json:"name" validate:"required"`
		Gender string `json:"gender" validate:"required,oneof=man woman"`
		Age    uint8  `json:"age" validate:"required,gte=0,lte=130"`
		Email  string `json:"email" validate:"required,email"`
	}

	// Updated user with valid values
	user := &User{
		ID:     1,
		Name:   "frank",
		Gender: "man", // Changed from "boy" to "man" to match the validation rule
		Age:    30,    // Changed from 180 to 30 to be within the valid range
		Email:  "gopher@88.com",
	}

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// Output the validation errors directly without translation
		for _, err := range validationErrors {
			fmt.Printf("Field '%s' failed validation on the '%s' tag\n", err.Field(), err.Tag())
		}
		return
	}

	fmt.Println("User validation passed!")
}
