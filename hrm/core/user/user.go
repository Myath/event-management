package user

import (
	"event-management/hrm/storage"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	Register(s storage.User) (*storage.User, error)
	GetUserByUsername(username string) (*storage.User, error)
	GetUserByID(id int) (*storage.User, error)
	UpdateUser(s storage.User) (*storage.User, error)
	DeleteUserByID(id int) error
	UserList(uf storage.UserFilter) ([]storage.User, error)
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

func (cu CoreUser) Login(s storage.Login) (*storage.User, error){
	user, err := cu.store.GetUserByUsername(s.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.Password)); err != nil{
		return nil, err
	}

	return user, nil
}

func (cu CoreUser) EditUser(s storage.User) (*storage.User, error){
	user, err := cu.store.GetUserByID(s.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cu CoreUser) UpdateUser(s storage.User) (*storage.User, error){
	hashPass, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	s.Password = string(hashPass)
	rUser, err := cu.store.UpdateUser(s)
	if err != nil {
		return nil, err
	}

	if rUser == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return rUser, nil
}

func (cu CoreUser) DeleteUser(s storage.User) error{
	if err := cu.store.DeleteUserByID(s.ID); err != nil {
		return nil
	}
	return nil
}

func (cu CoreUser) UserListWithFilter(s storage.UserFilter) ([]storage.User, error){
	uf := storage.UserFilter{
		SearchTerm: s.SearchTerm,
	}
	userList, err := cu.store.UserList(uf)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
