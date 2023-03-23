package user

import (
	"context"
	userpb "event-management/gunk/v1/user"
	"event-management/hrm/storage"
)

type UserCore interface {
	Register(s storage.User) (*storage.User, error)
	Login(s storage.Login) (*storage.User, error)
	EditUser(s storage.User) (*storage.User, error)
	UpdateUser(s storage.User) (*storage.User, error)
	DeleteUser(s storage.User) error
	UserListWithFilter(s storage.UserFilter) ([]storage.User, error)
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

func (us UserSvc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Username:  r.GetUsername(),
		Email:     r.GetEmail(),
		Password:  r.GetPassword(),
		IsAdmin: r.GetIsAdmin(),
	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.Core.Register(user)
	if err != nil {
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

func (us UserSvc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	loginUser := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := loginUser.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.Core.Login(loginUser)
	if err != nil {
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

func (us UserSvc) EditUser(ctx context.Context, r *userpb.EditUserRequest) (*userpb.EditUserResponse, error) {
	userID := storage.User{
		ID: int(r.GetID()),
	}

	u, err := us.Core.EditUser(userID)
	if err != nil {
		return nil, err
	}

	return &userpb.EditUserResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Password:  u.Password,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) UpdateUser(ctx context.Context, r *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := storage.User{
		ID:        int(r.GetID()),
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
		IsActive:  r.GetIsActive(),
	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.Core.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Password:  u.Password,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) DeleteUser(ctx context.Context, r *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	userID := storage.User{
		ID: int(r.GetID()),
	}

	err := us.Core.DeleteUser(userID)
	if err != nil {
		return nil, err
	}

	return &userpb.DeleteUserResponse{}, nil
}

func (us UserSvc) ListUser(ctx context.Context, r *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {
	userList := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
	}

	u, err := us.Core.UserListWithFilter(userList)
	if err != nil {
		return nil, err
	}

	var allUsers []*userpb.User
	for _, single := range u {
		user := &userpb.User{
			ID:        int32(single.ID),
			FirstName: single.FirstName,
			LastName:  single.LastName,
			Username:  single.Username,
			Email:     single.Email,
			Password:  single.Password,
			IsAdmin:   single.IsAdmin,
			IsActive:  single.IsActive,
		}
		allUsers = append(allUsers, user)
	}

	return &userpb.ListUserResponse{
		UserFilterList: &userpb.UserFilterList{
			AllUser:    allUsers,
			SearchTerm: userList.SearchTerm,
		},
	}, nil
}
