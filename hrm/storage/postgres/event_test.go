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

func TestInsertEvent(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenTypes := []storage.EventTypes{
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

	for _, evenType := range evenTypes {
		_, err := s.CreateEventType(evenType)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	event := storage.Event{
		EventTypeId: 1,
		EventName:   "Tour Of Sundarbon 2033",
		Description: "Tour of Sundarbon 2023, All are invited",
		Location:    "Sundarbon",
		StartAt:     time.Time{},
		EndAt:       time.Time{},
		PublishedAt: sql.NullTime{},
	}

	tests := []struct {
		name    string
		in      storage.Event
		want    *storage.Event
		wantErr bool
	}{
		{
			name: "INSERT_EVENT_SUCCESS",
			in:   event,
			want: &event,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.InsertEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.InsertEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Event{}, "ID", "EventTypeName", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.InsertEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetEventByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenTypes := []storage.EventTypes{
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

	for _, evenType := range evenTypes {
		_, err := s.CreateEventType(evenType)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	event := storage.Event{
		EventTypeId: 1,
		EventName:   "Tour Of Sundarbon 2033",
		Description: "Tour of Sundarbon 2023, All are invited",
		Location:    "Sundarbon",
		StartAt:     time.Time{},
		EndAt:       time.Time{},
		PublishedAt: sql.NullTime{},
	}

	createEvent, err := s.InsertEvent(event)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := createEvent.ID

	tests := []struct {
		name    string
		in      int
		want    *storage.Event
		wantErr bool
	}{
		{
			name: "GET_EVENT_BY_ID_SUCCESS",
			in:   id,
			want: &event,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetEventByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetEventByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Event{}, "ID", "EventTypeName", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.InsertEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateEvent(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenTypes := []storage.EventTypes{
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

	for _, evenType := range evenTypes {
		_, err := s.CreateEventType(evenType)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	event := storage.Event{
		EventTypeId: 1,
		EventName:   "Tour Of Sundarbon 2033",
		Description: "Tour of Sundarbon 2023, All are invited",
		Location:    "Sundarbon",
		StartAt:     time.Time{},
		EndAt:       time.Time{},
		PublishedAt: sql.NullTime{},
	}

	createEvent, err := s.InsertEvent(event)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertEvent() error = %v", err)
	}
	id := createEvent.ID

	tests := []struct {
		name    string
		in      storage.Event
		want    *storage.Event
		wantErr bool
	}{
		{
			name: "UPDATE_EVENT_SUCCESS",
			in: storage.Event{
				EventTypeId: 2,
				EventName:   "Tour Of Sundarbon",
				Description: "Tour of Sundarbon, All are invited",
				Location:    "Koromjol",
				StartAt:     time.Time{},
				EndAt:       time.Time{},
				PublishedAt: sql.NullTime{},
			},
			want: &storage.Event{
				ID:          id,
				EventTypeId: 2,
				EventName:   "Tour Of Sundarbon",
				Description: "Tour of Sundarbon, All are invited",
				Location:    "Koromjol",
				StartAt:     time.Time{},
				EndAt:       time.Time{},
				PublishedAt: sql.NullTime{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.in.ID = id
			got, err := s.UpdateEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Event{}, "EventTypeName", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteEventByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

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
			t.Fatalf("unable to create eventType for list event testing %v", err)
		}
	}

	event := storage.Event{
		EventTypeId: 1,
		EventName:   "Tour Of Sundarbon 2033",
		Description: "Tour of Sundarbon 2023, All are invited",
		Location:    "Sundarbon",
		StartAt:     time.Time{},
		EndAt:       time.Time{},
		PublishedAt: sql.NullTime{},
	}

	createEvent, err := s.InsertEvent(event)
	if err != nil {
		t.Fatalf("PostgresStorage.Event() error = %v", err)
	}
	id := createEvent.ID

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_EVENT_SUCCESS",
			in:   id,
		},
		{
			name:    "DELETE_EVENT_FAILED",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := s.DeleteEventByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteEventByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestListEvent(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

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
			EventTypeId: 3,
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

	tests := []struct {
		name    string
		in      storage.EventFilter
		want    []storage.Event
		wantErr bool
	}{
		{
			name: "EVENT_LIST_SUCCESS",
			in:   storage.EventFilter{},
			want: events,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.ListEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Event{}, "ID", "EventTypeName", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetPublishedDateByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenTypes := []storage.EventTypes{
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

	for _, evenType := range evenTypes {
		_, err := s.CreateEventType(evenType)
		if err != nil {
			t.Fatalf("unable to create eventType for list event testing %v", err)
		}
	}

	event := storage.Event{
		EventTypeId: 1,
		EventName:   "Tour Of Sundarbon 2033",
		Description: "Tour of Sundarbon 2023, All are invited",
		Location:    "Sundarbon",
		StartAt:     time.Time{},
		EndAt:       time.Time{},
		PublishedAt: sql.NullTime{},
	}

	createEvent, err := s.InsertEvent(event)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertEvent() error = %v", err)
	}
	id := createEvent.ID

	tests := []struct {
		name    string
		in      int
		want    *storage.Event
		wantErr bool
	}{
		{
			name: "GET_PUBLISHED_DATE_FOR_PUBLISHED_FORM_SUCCESS",
			in:   id,
			want: &storage.Event{
				ID:        id,
				DeletedAt: sql.NullTime{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetPublishedDateByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetPublishedDateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("PostgresStorage.InsertEvent() diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestPublishedEvent(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenTypes := []storage.EventTypes{
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

	for _, evenType := range evenTypes {
		_, err := s.CreateEventType(evenType)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	event := storage.Event{
		EventTypeId: 1,
		EventName:   "Tour Of Sundarbon 2033",
		Description: "Tour of Sundarbon 2023, All are invited",
		Location:    "Sundarbon",
		StartAt:     time.Time{},
		EndAt:       time.Time{},
		PublishedAt: sql.NullTime{},
	}

	createEvent, err := s.InsertEvent(event)
	if err != nil {
		t.Fatalf("PostgresStorage.InsertEvent() error = %v", err)
	}
	id := createEvent.ID

	tests := []struct {
		name    string
		in      storage.Event
		want    *storage.Event
		wantErr bool
	}{
		{
			name: "UPDATE_EVENT_SUCCESS",
			in: storage.Event{
				PublishedAt: sql.NullTime{},
			},
			want: &storage.Event{
				ID:          id,
				EventTypeId: 1,
				EventName:   "Tour Of Sundarbon 2033",
				Description: "Tour of Sundarbon 2023, All are invited",
				Location:    "Sundarbon",
				StartAt:     time.Time{},
				EndAt:       time.Time{},
				PublishedAt: sql.NullTime{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.in.ID = id
			got, err := s.PublishedEvent(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.PublishedEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Event{}, "EventTypeName", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateEvent() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
