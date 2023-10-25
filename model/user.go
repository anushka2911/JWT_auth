package model

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	if err := validate.RegisterValidation("custom_validate_password", customValidatePassword); err != nil {
		fmt.Printf("Failed to register custom validation: %v\n", err)
	}
	fmt.Println("Successfully registered custom validation")
}

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	UserID       string `json:"user_id"`
	UserType     string `json:"user_type" validate:"required,min=1,max=20"`
	FirstName    string `json:"first_name" validate:"required,min=1,max=20"`
	LastName     string `json:"last_name" validate:"required,min=1,max=20"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=1,max=10"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6,max=20,custom_validate_password"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

func customValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasAtleastOneCapitalLetter := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasAtleastOneSmallLetter := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasAtleastOneNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasAtleastOneSpecialCharacter := regexp.MustCompile(`[!@#$&*]`).MatchString(password)

	return hasAtleastOneCapitalLetter && hasAtleastOneSmallLetter && hasAtleastOneNumber && hasAtleastOneSpecialCharacter
}

func (u *User) ValidateUserInput() error {
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}
