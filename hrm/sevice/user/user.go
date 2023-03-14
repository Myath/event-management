package user

import (
	"context"
	userpb "event-management/gunk/v1/user"
	"event-management/hrm/storage"
)

type UserCore interface {
	Register(s storage.User) (*storage.User, error)
	Login(s storage.Login) (*storage.User, error)
}

type UserSvc struct {
	userpb.UnimplementedUserServiceServer
	Core UserCore
}

func NewUserSvc(uc UserCore) *UserSvc {
	return &UserSvc{
		Core: uc,
	}
}

func (us UserSvc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error){
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName: r.GetLastName(),
		Username: r.GetUsername(),
		Email: r.GetEmail(),
		Password: r.GetPassword(),
	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.Core.Register(user)
	if err != nil{
		return nil, err
	}

	return &userpb.RegisterResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Password:  u.Password,
			IsAdmin:   u.IsAdmin,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error){
	loginUser := storage.Login{
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
	}

	if err := loginUser.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.Core.Login(loginUser)
	if err != nil{
		return nil, err
	}

	return &userpb.LoginResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Password:  u.Password,
			IsAdmin:   u.IsAdmin,
			IsActive:  u.IsActive,
		},
	}, nil
}