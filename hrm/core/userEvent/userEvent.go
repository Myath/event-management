package userevent

import (
	"event-management/hrm/storage"
	"fmt"
)
	

type UserEventStore interface {
	InsertUserEvent(s storage.UserEvent) (*storage.UserEvent, error)
	GetUserEventByIDQuery(userId, eventId int) (*storage.UserEvent, error)
	UpdateUserEvent(s storage.UserEvent) (*storage.UserEvent, error)
}

type CoreUserEvent struct {
	store UserEventStore
}

func NewCoreUserEvent (us UserEventStore) *CoreUserEvent{
	return &CoreUserEvent{
		store: us,
	}
}

func (cue CoreUserEvent) InsertUserEvent(s storage.UserEvent) (*storage.UserEvent, error) {
	userEvent,  err := cue.store.InsertUserEvent(s)
	if err != nil {
		return nil, err
	}

	if userEvent == nil {
		return nil, fmt.Errorf("unable to insert userEvent")
	}

	return userEvent, nil
}


func (cue CoreUserEvent) IsEventAvailable(userID, eventId int) (*storage.UserEvent, error) {
	userEvent,  err := cue.store.GetUserEventByIDQuery(userID, eventId)
	if err != nil {
		return nil, err
	}

	if userEvent == nil {
		return nil, fmt.Errorf("unable to insert userEvent")
	}

	return userEvent, nil
}

func (cue CoreUserEvent) UpdateUserEvent(s storage.UserEvent) (*storage.UserEvent, error) {
	userEvent,  err := cue.store.UpdateUserEvent(s)
	if err != nil {
		return nil, err
	}

	if userEvent == nil {
		return nil, fmt.Errorf("unable to insert userEvent")
	}

	return userEvent, nil
}