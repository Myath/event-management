package postgres

import (
	"database/sql"
	"event-management/hrm/storage"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestInsertComment(t *testing.T) {
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

	comment := storage.Comments{
		UserId:  1,
		EventID: 2,
		Comment: "I want to join this Tour",
	}

	tests := []struct {
		name    string
		in      storage.Comments
		want    *storage.Comments
		wantErr bool
	}{
		{
			name: "CREATE_COMMENT_SUCCESS",
			in:   comment,
			want: &comment,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.InsertComment(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.InsertComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Comments{}, "ID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.InsertComment() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetCommentID(t *testing.T) {
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

	comment := storage.Comments{
		UserId:  1,
		EventID: 2,
		Comment: "I want to join this Tour",
	}
	createComment, err := s.InsertComment(comment)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertComment() error = %v", err)
	}
	id := createComment.ID

	tests := []struct {
		name    string
		in      int
		want    *storage.Comments
		wantErr bool
	}{
		{
			name: "GET_COMMENT_BY_ID_SUCCESS",
			in:   id,
			want: &comment,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetCommentID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCommentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Comments{}, "ID", "CreatedAt", "UpdatedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetCommentID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateComment(t *testing.T) {
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

	comment := storage.Comments{
		UserId:  1,
		EventID: 2,
		Comment: "I want to join this Tour",
	}
	createComment, err := s.InsertComment(comment)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertComment() error = %v", err)
	}
	id := createComment.ID

	tests := []struct {
		name    string
		in      storage.Comments
		want    *storage.Comments
		wantErr bool
	}{
		{
			name: "UPDATE_COMMENT_SUCCESS",
			in: storage.Comments{
				Comment: "I want to cancel this Tour",
			},
			want: &storage.Comments{
				UserId:  1,
				EventID: 2,
				Comment: "I want to cancel this Tour",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = id
			got, err := s.UpdateComment(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Comments{}, "ID", "CreatedAt", "UpdatedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateComment() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestCommentsOfEvent(t *testing.T) {
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
	comments := []storage.Comments{
		{
			UserId:  1,
			EventID: 2,
			Comment: "I want to join this Tour",
		},

		{
			UserId:  2,
			EventID: 2,
			Comment: "I want to join this Tour",
		},

		{
			UserId:  3,
			EventID: 2,
			Comment: "I want to join this Tour",
		},
	}

	for _, comment := range comments {
		_, err := s.InsertComment(comment)
		if err != nil {
			t.Fatalf("unable to create comment for list comment testing %v", err)
		}
	}

	sort.SliceStable(comments, func(i, j int) bool {
		return comments[i].UserId > comments[j].UserId
	})

	tests := []struct {
		name    string
		in      int
		want    []storage.Comments
		wantErr bool
	}{
		{
			name: "EVENT_LIST_SUCCESS",
			in:   2,
			want: comments,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CommentsOfEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CommentsOfEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Comments{}, "ID","Username", "IsAdmin", "EventName", "EventTypeId", "EventTypeName", "Description", "StartAt", "EndAt", "Location",  "CreatedAt", "UpdatedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})
			
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CommentsOfEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
