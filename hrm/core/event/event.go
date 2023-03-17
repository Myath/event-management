package event

import (
	"event-management/hrm/storage"
	"fmt"
)


type EventStore interface {
	InsertEvent(s storage.Event) (*storage.Event, error)
	GetEventByID(id int) (*storage.Event, error)
	UpdateEvent(s storage.Event) (*storage.Event, error)
}

type CoreEvent struct {
	store EventStore
}

func NewCoreEvent (es EventStore) *CoreEvent{
	return &CoreEvent{
		store: es,
	}
}

func (ce CoreEvent) InsertEvent(s storage.Event) (*storage.Event, error){

	event , err := ce.store.InsertEvent(s)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, fmt.Errorf("unable to insert event")
	}

	return event, nil
}

func (ce CoreEvent) EditUser(s storage.Event) (*storage.Event, error){
	event, err := ce.store.GetEventByID(s.ID)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (ce CoreEvent) UpdateEvent(s storage.Event) (*storage.Event, error){
	event, err := ce.store.UpdateEvent(s)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, fmt.Errorf("unable to update event")
	}

	return event, nil
}