package entity

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-jwt/jwt/v5"
)

type RegistrationPayload struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	Id        string `db:"id"`
	Email     string `db:"email"`
	Name      string `db:"name"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTPayload struct {
	Id    string
	Email string
	Name  string
}

type JWTClaims struct {
	Id    string
	Email string
	Name  string
	jwt.RegisteredClaims
}

func NewUser(email, name, password string) *User {
	u := &User{
		Email:    email,
		Name:     name,
		Password: password,
	}

	return u
}

func (u *Credential) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Email,
			validation.Required.Error("email is required"),
			validation.By(validateEmailFormat),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password is required"),
			validation.Length(5, 15).Error("password must be between 5 and 15 characters"),
		),
	)

	return err
}

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Email,
			validation.Required.Error("email is required"),
			validation.By(validateEmailFormat),
		),
		validation.Field(&u.Name,
			validation.Required.Error("name is required"),
			validation.Length(5, 50).Error("name must be between 5 and 50 characters"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password is required"),
			validation.Length(5, 15).Error("password must be between 5 and 15 characters"),
		),
	)

	return err
}

func validateEmailFormat(value any) error {
	email, ok := value.(string)
	if !ok {
		return errors.New("parse error")
	}

	pattern := "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	rgx := regexp.MustCompile(pattern)
	if !rgx.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}
