package eventType

import (
	"event-management/hrm/storage"
	"fmt"
)

type EventTypeStore interface {
	CreateEventType(s storage.EventTypes) (*storage.EventTypes, error)
	GetEventTypeByID(id int) (*storage.EventTypes, error)
	UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error)
	DeleteEventTypeByID(id int) error
	EventTypeList(uf storage.EventTypesFilter) ([]storage.EventTypes, error)
}

type CoreEventType struct {
	store EventTypeStore
}

func NewCoreEventType (es EventTypeStore) *CoreEventType{
	return &CoreEventType{
		store: es,
	}
}

func (ce CoreEventType) CreateEventType(s storage.EventTypes) (*storage.EventTypes, error){
	eventType, err := ce.store.CreateEventType(s)
	if err != nil {
		return nil, err
	}

	if eventType == nil {
		return nil, fmt.Errorf("unable to createEvnetType")
	}

	return eventType, nil
}

func (ce CoreEventType) EditEventType(s storage.EventTypes) (*storage.EventTypes, error){
	eventType, err := ce.store.GetEventTypeByID(s.ID)
	if err != nil {
		return nil, err
	}
	return eventType, nil

}

func (ce CoreEventType) UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error){
	eventType, err := ce.store.UpdateEventType(s)
	if err != nil {
		return nil, err
	}

	if eventType == nil {
		return nil, fmt.Errorf("unable to update eventType")
	}

	return eventType, nil
}

func (ce CoreEventType) DeleteEventType(s storage.EventTypes) error{
	if err := ce.store.DeleteEventTypeByID(s.ID); err != nil {
		return nil
	}
	return nil
}

func (ce CoreEventType) EventTypeListWithFilter (s storage.EventTypesFilter) ([]storage.EventTypes, error){
	uf := storage.EventTypesFilter{
		SearchTerm: s.SearchTerm,
	}

	eventTypeList, err := ce.store.EventTypeList(uf)
	if err != nil {
		return nil, err
	}

	return eventTypeList, nil
}