package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID        int          `json:"id" form:"-" db:"id"`
	FirstName string       `json:"first_name" db:"first_name"`
	LastName  string       `json:"last_name" db:"last_name"`
	Email     string       `json:"email" db:"email"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	IsAdmin   bool         `json:"is_admin" db:"is_admin"`
	IsActive  bool         `json:"is_active" db:"is_active"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.Username,
			validation.Required.When(u.ID == 0).Error("The username field is required."),
		),
		validation.Field(&u.Email,
			validation.Required.When(u.ID == 0).Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.When(u.ID == 0).Error("The password field is required."),
		),
	)
}

type Login struct {
	Username string
	Password string
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&l.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}
