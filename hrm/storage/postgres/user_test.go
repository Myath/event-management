package postgres

import (
	"event-management/hrm/storage"
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRegister(t *testing.T) {

	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Rahim",
		LastName:  "Sheikh",
		Username:  "rahim",
		Email:     "rahimsehikh@gmail.com",
		Password:  "Rahim@143",
		IsAdmin:   false,
		IsActive:  true,
	}

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "USER_REGISTER_SUCCESS",
			in:   user,
			want: &storage.User{
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim",
				Email:     "rahimsehikh@gmail.com",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
		{
			name: "REGISTER_DUPLICATE_USERNAME_FAILURE",
			in: storage.User{
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim",
				Email:     "rahimsehikh1@gmail.com",
				Password:  "Rahim@143",
				IsAdmin:   false,
				IsActive:  true,
			},
			wantErr: true,
		},
		{
			name: "REGISTER_DUPLICATE_EMAIL_FAILURE",
			in: storage.User{
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim2",
				Email:     "rahimsehikh@gmail.com",
				Password:  "Rahim@143",
				IsAdmin:   false,
				IsActive:  true,
			},
			wantErr: true,
		},

		{
			name: "ADMIN_REGISTER_SUCCESS",
			in: storage.User{
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim675",
				Email:     "rahimsehikh675@gmail.com",
				Password:  "Rahim@143",
				IsAdmin:   true,
				IsActive:  true,
			},
			want: &storage.User{
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim675",
				Email:     "rahimsehikh675@gmail.com",
				IsAdmin:   true,
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.in.FirstName)
			got, err := s.Register(tt.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Rahim",
		LastName:  "Sheikh",
		Username:  "rahim",
		Email:     "rahimsehikh@gmail.com",
		Password:  "Rahim@143",
		IsAdmin:   false,
		IsActive:  true,
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	username := createUser.Username

	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET_USER_BY_USERNAME_SUCCESS",
			in:   username,
			want: &storage.User{
				ID:        createUser.ID,
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim",
				Email:     "rahimsehikh@gmail.com",
				Password:  "Rahim@143",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserByUsername(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetUserByUsername() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Rahim",
		LastName:  "Sheikh",
		Username:  "rahim",
		Email:     "rahimsehikh@gmail.com",
		Password:  "Rahim@143",
		IsAdmin:   false,
		IsActive:  true,
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := createUser.ID

	tests := []struct {
		name    string
		in      int
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET_USER_BY_ID_SUCCESS",
			in:   id,
			want: &storage.User{
				ID:        id,
				FirstName: "Rahim",
				LastName:  "Sheikh",
				Username:  "rahim",
				Email:     "rahimsehikh@gmail.com",
				Password:  "Rahim@143",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetUserByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Rahim",
		LastName:  "Sheikh",
		Username:  "rahim",
		Email:     "rahimsehikh@gmail.com",
		Password:  "Rahim@143",
		IsAdmin:   false,
		IsActive:  true,
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	id := createUser.ID

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "USER_UPDATE_SUCCESS",
			in: storage.User{
				FirstName: "Josim",
				LastName:  "Uddin",
			},
			want: &storage.User{
				ID:        id,
				FirstName: "Josim",
				LastName:  "Uddin",
				Username:  "rahim",
				Email:     "rahimsehikh@gmail.com",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = id
			got, err := s.UpdateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteUserByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Rahim",
		LastName:  "Sheikh",
		Username:  "rahim",
		Email:     "rahimsehikh@gmail.com",
		Password:  "Rahim@143",
		IsAdmin:   false,
		IsActive:  true,
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	id := createUser.ID

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_USER_BY_ID_SUCCESS",
			in:   id,
		},
		{
			name:    "DELETE_USER_BY_ID_FAILED",
			in:      id,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteUserByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserList(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	users := []storage.User{
		{
			FirstName: "Rahim",
			LastName:  "Sheikh",
			Username:  "rahim",
			Email:     "rahimsehikh@gmail.com",
			Password:  "Rahim@143",
			IsAdmin:   false,
			IsActive:  true,
		},
		{
			FirstName: "Jony",
			LastName:  "Ghose",
			Username:  "jony",
			Email:     "jony@gmail.com",
			Password:  "Rahim@143",
			IsAdmin:   false,
			IsActive:  true,
		},
		{
			FirstName: "Apu",
			LastName:  "Sarder",
			Username:  "apu",
			Email:     "apu@gmail.com",
			Password:  "Rahim@143",
			IsAdmin:   false,
			IsActive:  true,
		},
	}

	for _, user := range users {
		_, err := s.Register(user)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.UserFilter
		want    []storage.User
		wantErr bool
	}{
		{
			name: "USER_LIST_SUCCESS",
			in:   storage.UserFilter{},
			want: users,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.UserList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
