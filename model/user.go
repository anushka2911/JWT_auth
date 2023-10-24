package model

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	UserID       string `json:"user_id"`
	FirstName    string `json:"first_name" validate:"required,min=1,max=20"`
	LastName     string `json:"last_name" validate:"required,min=1,max=20"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=1,max=10"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6,max=20 custom_validate_password"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

func customValidatePassword(password validator.FieldLevel) bool {
	hasAtleastOneCapitalLetter := regexp.MustCompile(`[A-Z]`).MatchString(password.Field().String())
	hasAtleastOneSmallLetter := regexp.MustCompile(`[a-z]`).MatchString(password.Field().String())
	hasAtleastOneNumber := regexp.MustCompile(`[0-9]`).MatchString(password.Field().String())
	hasAtleastOneSpecialCharacter := regexp.MustCompile(`[!@#$&*]`).MatchString(password.Field().String())

	return hasAtleastOneCapitalLetter && hasAtleastOneSmallLetter && hasAtleastOneNumber && hasAtleastOneSpecialCharacter
}

func (u *User) Validate() error {

	err := validate.RegisterValidation("custom_validate_password", customValidatePassword)
	if err != nil {
		return err
	}

	return validate.Struct(u)
}
