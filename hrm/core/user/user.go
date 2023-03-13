package user

import (
	"event-management/hrm/storage"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	Register(s storage.User) (*storage.User, error)
}

type CoreUser struct{
	store UserStore
}

func NewCoreUser(us UserStore) *CoreUser {
	return &CoreUser{
		store: us,
	}
}

func (cu CoreUser) Register(s storage.User) (*storage.User, error){
	hashPass, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	s.Password = string(hashPass)
	rUser, err := cu.store.Register(s)
	if err != nil {
		return nil, err
	}

	if rUser == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return rUser, nil
}