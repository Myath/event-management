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
	Username  string       `json:"username" db:"username"`
	Email     string       `json:"email" db:"email"`
	Password  string       `json:"password" db:"password"`
	IsAdmin   bool         `json:"is_admin" db:"is_admin"`
	IsActive  bool         `json:"is_active" db:"is_active"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type UserFilter struct {
	SearchTerm string
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

type EventTypes struct {
	ID            int          `json:"id" form:"-" db:"id"`
	EventTypeName string       `json:"event_type_name" db:"event_type_name"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type EventTypesFilter struct {
	SearchTerm string
}

func (e EventTypes) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.EventTypeName,
			validation.Required.Error("The event name is required."),
		),
	)
}

type Event struct {
	ID            int          `json:"id" form:"-" db:"id"`
	EventTypeId   int          `json:"event_type_id" db:"event_type_id"`
	EventTypeName string       `json:"event_type_name" db:"event_type_name"`
	EventName     string       `json:"event_name" db:"event_name"`
	Description   string       `json:"description" db:"description"`
	Location      string       `json:"location" db:"location"`
	StartAt       time.Time    `json:"start_at" db:"start_at"`
	EndAt         time.Time    `json:"end_at" db:"end_at"`
	PublishedAt   sql.NullTime `json:"published_at" db:"published_at"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type EventFilter struct {
	SearchTerm string
}

func (e Event) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.EventTypeId,
			validation.Required.Error("The eventTypeId is required."),
		),
		validation.Field(&e.EventName,
			validation.Required.Error("The event name field is required."),
		),
		validation.Field(&e.Description,
			validation.Required.When(e.ID == 0).Error("The description field is required."),
		),
		validation.Field(&e.Location,
			validation.Required.When(e.ID == 0).Error("The location field is required."),
		),
		validation.Field(&e.StartAt,
			validation.Required.When(e.ID == 0).Error("The start time field is required."),
		),
		validation.Field(&e.EndAt,
			validation.Required.When(e.ID == 0).Error("The end time field is required."),
		),
	)
}

type Comments struct {
	ID            int          `json:"id" form:"-" db:"id"`
	UserId        int          `json:"user_id" db:"user_id"`
	Username      string       `json:"username" db:"username"`
	IsAdmin       bool         `json:"is_admin" db:"is_admin"`
	EventID       int          `json:"event_id" db:"event_id"`
	EventName     string       `json:"event_name" db:"event_name"`
	EventTypeId   int          `json:"event_type_id" db:"event_type_id"`
	EventTypeName string       `json:"event_type_name" db:"event_type_name"`
	Description   string       `json:"description" db:"description"`
	Location      string       `json:"location" db:"location"`
	StartAt       time.Time    `json:"start_at" db:"start_at"`
	EndAt         time.Time    `json:"end_at" db:"end_at"`
	PublishedAt   sql.NullTime `json:"published_at" db:"published_at"`
	Comment       string       `json:"comment" db:"comment"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
}

func (c Comments) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.UserId,
			validation.Required.Error("The userId is required."),
		),
		validation.Field(&c.EventID,
			validation.Required.Error("The eventID name field is required."),
		),
		validation.Field(&c.Comment,
			validation.Required.When(c.ID == 0).Error("The Comment field is required."),
		),
	)
}

type UserEvent struct {
	UserId    int       `json:"user_id" db:"user_id"`
	EventID   int       `json:"event_id" db:"event_id"`
	Status    int       `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u UserEvent) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserId,
			validation.Required.Error("The userId is required."),
		),
		validation.Field(&u.EventID,
			validation.Required.Error("The eventID name field is required."),
		),
	)
}
