package postgres

import (
	"database/sql"
	"event-management/hrm/storage"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestPostgresStorage_InsertUserEvent(t *testing.T) {
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

	eventTypes := []storage.EventTypes{
		{
			EventTypeName: "Tour",
		},

		{
			EventTypeName: "wedding",
		},

		{
			EventTypeName: "Catering",
		},
	}

	for _, eventType := range eventTypes {
		_, err := s.CreateEventType(eventType)
		if err != nil {
			t.Fatalf("unable to create evenType for list eventType testing %v", err)
		}
	}

	events := []storage.Event{
		{
			EventTypeId: 1,
			EventName:   "Tour Of Sundarbon 2033",
			Description: "Tour of Sundarbon 2023, All are invited",
			Location:    "Sundarbon",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
		{
			EventTypeId: 2,
			EventName:   "Tour Of Kuyakta 2033",
			Description: "Tour of Kuyakata 2023, All are invited",
			Location:    "Kuyakata",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
		{
			EventTypeId: 1,
			EventName:   "Happy Wedding of Shima+Rakib",
			Description: "Happy Wedding of Shima+Rakib, All are invited",
			Location:    "Khulna",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
	}

	for _, event := range events {
		_, err := s.InsertEvent(event)
		if err != nil {
			t.Fatalf("unable to create event for list eventType testing %v", err)
		}
	}

	userEvent := storage.UserEvent{
		UserId:  2,
		EventID: 1,
		Status:  0,
	}
	tests := []struct {
		name    string
		in      storage.UserEvent
		want    *storage.UserEvent
		wantErr bool
	}{
		{
			name: "CREATE_COMMENT_SUCCESS",
			in:   userEvent,
			want: &userEvent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.InsertUserEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.InsertUserEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.UserEvent{}, "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.InsertUserEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetUserEventByIDQuery(t *testing.T) {
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

	eventTypes := []storage.EventTypes{
		{
			EventTypeName: "Tour",
		},

		{
			EventTypeName: "wedding",
		},

		{
			EventTypeName: "Catering",
		},
	}

	for _, eventType := range eventTypes {
		_, err := s.CreateEventType(eventType)
		if err != nil {
			t.Fatalf("unable to create evenType for list eventType testing %v", err)
		}
	}

	events := []storage.Event{
		{
			EventTypeId: 1,
			EventName:   "Tour Of Sundarbon 2033",
			Description: "Tour of Sundarbon 2023, All are invited",
			Location:    "Sundarbon",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
		{
			EventTypeId: 2,
			EventName:   "Tour Of Kuyakta 2033",
			Description: "Tour of Kuyakata 2023, All are invited",
			Location:    "Kuyakata",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
		{
			EventTypeId: 1,
			EventName:   "Happy Wedding of Shima+Rakib",
			Description: "Happy Wedding of Shima+Rakib, All are invited",
			Location:    "Khulna",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
	}

	for _, event := range events {
		_, err := s.InsertEvent(event)
		if err != nil {
			t.Fatalf("unable to create event for list eventType testing %v", err)
		}
	}

	userEvent := storage.UserEvent{
		UserId:  2,
		EventID: 1,
		Status:  0,
	}

	createUserEvent, err := s.InsertUserEvent(userEvent)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertUserEvent() error = %v", err)
	}
	userID := createUserEvent.UserId
	eventID := createUserEvent.EventID
	tests := []struct {
		name    string
		in      storage.UserEvent
		want    *storage.UserEvent
		wantErr bool
	}{
		{
			name: "GET_USER_EVENT_BY_ID_SUCCESS",
			in: storage.UserEvent{
				UserId:  userID,
				EventID: eventID,
			},
			want: &userEvent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserEventByIDQuery(tt.in.UserId, tt.in.EventID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserEventByIDQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.UserEvent{}, "CreatedAt", "UpdatedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetUserEventByIDQuery() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateUserEvent(t *testing.T) {
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

	eventTypes := []storage.EventTypes{
		{
			EventTypeName: "Tour",
		},

		{
			EventTypeName: "wedding",
		},

		{
			EventTypeName: "Catering",
		},
	}

	for _, eventType := range eventTypes {
		_, err := s.CreateEventType(eventType)
		if err != nil {
			t.Fatalf("unable to create evenType for list eventType testing %v", err)
		}
	}

	events := []storage.Event{
		{
			EventTypeId: 1,
			EventName:   "Tour Of Sundarbon 2033",
			Description: "Tour of Sundarbon 2023, All are invited",
			Location:    "Sundarbon",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
		{
			EventTypeId: 2,
			EventName:   "Tour Of Kuyakta 2033",
			Description: "Tour of Kuyakata 2023, All are invited",
			Location:    "Kuyakata",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
		{
			EventTypeId: 1,
			EventName:   "Happy Wedding of Shima+Rakib",
			Description: "Happy Wedding of Shima+Rakib, All are invited",
			Location:    "Khulna",
			StartAt:     time.Time{},
			EndAt:       time.Time{},
			PublishedAt: sql.NullTime{},
		},
	}

	for _, event := range events {
		_, err := s.InsertEvent(event)
		if err != nil {
			t.Fatalf("unable to create event for list eventType testing %v", err)
		}
	}

	userEvent := storage.UserEvent{
		UserId:  2,
		EventID: 1,
		Status:  0,
	}

	createUserEvent, err := s.InsertUserEvent(userEvent)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertUserEvent() error = %v", err)
	}
	userID := createUserEvent.UserId
	eventID := createUserEvent.EventID

	tests := []struct {
		name    string
		in      storage.UserEvent
		want    *storage.UserEvent
		wantErr bool
	}{
		{
			name: "USER_EVENT_UPDATE_SUCCESS",
			in: storage.UserEvent{
				Status:    1,
			},
			want: &storage.UserEvent{
				UserId:    userID,
				EventID:   eventID,
				Status:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.UserId = userID
			tt.in.EventID = eventID
			got, err := s.UpdateUserEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUserEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.UserEvent{}, "CreatedAt", "UpdatedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUserEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
