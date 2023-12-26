package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `validate:"required"`
	Age   int    `validate:"required,gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func main() {
	user := &User{
		Age:   135,
		Email: "nas.com",
	}
	errMap := make(map[string]string)
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		tagName := err.Tag()

		var errMsg string
		switch tagName {
		case "required":
			errMsg = fmt.Sprintf("%s is required", fieldName)
		case "email":
			errMsg = fmt.Sprintf("%s must be a valid email address", fieldName)
		case "gte":
			errMsg = fmt.Sprintf("%s must be greater than or equal to %s", fieldName, err.Param())
		case "lte":
			errMsg = fmt.Sprintf("%s must be greater than or equal to %s", fieldName, err.Param())
		default:
			errMsg = fmt.Sprintf("%s us invalid", fieldName)
		}
		errMap[fieldName] = errMsg
	}
	fmt.Println(errMap)
}
