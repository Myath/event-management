package event

import (
	"event-management/hrm/storage"
	"fmt"
)

type EventStore interface {
	CreateEventType(s storage.EventTypes) (*storage.EventTypes, error)
	GetEventTypeByID(id int) (*storage.EventTypes, error)
	UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error)
}

type CoreEvent struct {
	store EventStore
}

func NewCoreEvent (es EventStore) *CoreEvent{
	return &CoreEvent{
		store: es,
	}
}

func (ce CoreEvent) CreateEventType(s storage.EventTypes) (*storage.EventTypes, error){
	eventType, err := ce.store.CreateEventType(s)
	if err != nil {
		return nil, err
	}

	if eventType == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return eventType, nil
}

func (ce CoreEvent) EditEventType(s storage.EventTypes) (*storage.EventTypes, error){
	eventType, err := ce.store.GetEventTypeByID(s.ID)
	if err != nil {
		return nil, err
	}
	return eventType, nil

}

func (ce CoreEvent) UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error){
	eventType, err := ce.store.UpdateEventType(s)
	if err != nil {
		return nil, err
	}

	if eventType == nil {
		return nil, fmt.Errorf("unable to update eventType")
	}

	return eventType, nil
}