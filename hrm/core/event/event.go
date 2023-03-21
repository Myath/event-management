package event

import (
	"event-management/hrm/storage"
	"fmt"
)


type EventStore interface {
	InsertEvent(s storage.Event) (*storage.Event, error)
	GetEventByID(id int) (*storage.Event, error)
	UpdateEvent(s storage.Event) (*storage.Event, error)
	DeleteEventByID(id int) error
	ListEvent(uf storage.EventFilter) ([]storage.Event, error)
	GetPublishedDateByID(id int) (*storage.Event, error)
	PublishedEvent(s storage.Event) (*storage.Event, error)
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

func (ce CoreEvent) DeleteEvent(s storage.Event) error{
	if err := ce.store.DeleteEventByID(s.ID); err != nil {
		return nil
	}
	return nil
}

 func (ce CoreEvent) EventListWithFilter(s storage.EventFilter) ([]storage.Event, error){
	uf := storage.EventFilter{
		SearchTerm: s.SearchTerm,
	}
	eventList, err := ce.store.ListEvent(uf)
	if err != nil {
		return nil, err
	}
	return eventList, nil
 }


 func (ce CoreEvent) PublishedDateForm(s storage.Event) (*storage.Event, error){
	publishedDate, err := ce.store.GetPublishedDateByID(s.ID)
	if err != nil {
		return nil, err
	}

	return publishedDate, nil
}


func (ce CoreEvent) PublishedEvent(s storage.Event) (*storage.Event, error){
	eventPublished, err := ce.store.PublishedEvent(s)
	if err != nil {
		return nil, err
	}

	if eventPublished == nil {
		return nil, fmt.Errorf("unable to update event")
	}

	return eventPublished, nil
}
