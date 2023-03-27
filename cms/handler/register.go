package handler

import (
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	userpb "event-management/gunk/v1/user"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"
)

type RegisterUser struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	IsAdmin   bool
	IsActive  bool
	FormError map[string]error
}

type RegisterUserFilter struct {
	SearchTerm string
}

func (u RegisterUser) Validate() error {
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

type RegisterUserForm struct {
	RegisterUser RegisterUser
	FormError    map[string]error
	CSRFToken    string
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	if h.SessionIDGet(w, r) != "" {
		http.Redirect(w, r, "/eventtype", http.StatusSeeOther)
		return
	}
	h.pareseRegisterTemplate(w, RegisterUserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) pareseRegisterTemplate(w http.ResponseWriter, form RegisterUserForm) {
	t := h.Templates.Lookup("registerUser.html")
	if t == nil {
		log.Println("unable to lookup register template")
		h.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, form); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) RegisterPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	user := RegisterUser{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			user.FormError = vErr
		}
		h.pareseRegisterTemplate(w, RegisterUserForm{
			RegisterUser: user,
			FormError:    user.FormError,
			CSRFToken:    nosurf.Token(r),
		})
		return
	}

	_, err := h.hrmSvc.Register(r.Context(), &userpb.RegisterRequest{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h Handler) SessionIDGet(w http.ResponseWriter, r *http.Request) string {
	userID := h.sessionManager.GetString(r.Context(), "userID")
	if userID == "" {
		return ""
	}
	return userID
}