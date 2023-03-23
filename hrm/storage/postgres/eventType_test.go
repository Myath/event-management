package postgres

import (
	"event-management/hrm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateEventType(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenType := storage.EventTypes{
		EventTypeName: "Tour",
	}

	tests := []struct {
		name    string
		in      storage.EventTypes
		want    *storage.EventTypes
		wantErr bool
	}{
		{
			name: "CREATE_EVENT_TYPE_SUCCESS",
			in:   evenType,
			want: &storage.EventTypes{
				EventTypeName: "Tour",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.CreateEventType(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateEventType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.EventTypes{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateEventType() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetEventTypeByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenType := storage.EventTypes{
		EventTypeName: "Tour",
	}

	createEventType, err := s.CreateEventType(evenType)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := createEventType.ID

	tests := []struct {
		name    string
		in      int
		want    *storage.EventTypes
		wantErr bool
	}{
		{
			name: "GET_EVENT_TYPE_BY_ID_SUCCESS",
			in:   id,
			want: &storage.EventTypes{
				ID:            id,
				EventTypeName: "Tour",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetEventTypeByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetEventTypeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.EventTypes{}, "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetEventTypeByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateEventType(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenType := storage.EventTypes{
		EventTypeName: "Tour",
	}

	createEventType, err := s.CreateEventType(evenType)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := createEventType.ID

	tests := []struct {
		name    string
		in      storage.EventTypes
		want    *storage.EventTypes
		wantErr bool
	}{
		{
			name: "UPDATE_EVENT_TYPE_SUCCESS",
			in: storage.EventTypes{
				EventTypeName: "Dhaka Travel Mart",
			},
			want: &storage.EventTypes{
				ID:            id,
				EventTypeName: "Dhaka Travel Mart",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = id
			got, err := s.UpdateEventType(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateEventType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.EventTypes{}, "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateEventType() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteEventTypeByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()
	t.Cleanup(func() {
		tr()
	})

	evenType := storage.EventTypes{
		EventTypeName: "Tour",
	}

	createEventType, err := s.CreateEventType(evenType)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := createEventType.ID

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_EVENT_TYPE_BY_ID_SUCCESS",
			in:   id,
		},
		{
			name:    "DELETE_EVENT_TYPE_BY_ID_FAILED",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := s.DeleteEventTypeByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteEventTypeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEventTypeList(t *testing.T) {
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

	tests := []struct {
		name    string
		in      storage.EventTypesFilter
		want    []storage.EventTypes
		wantErr bool
	}{
		{
			name: "USER_LIST_SUCCESS",
			in:   storage.EventTypesFilter{},
			want: evenTypes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.EventTypeList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.EventTypeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.EventTypes{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.EventTypeList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
