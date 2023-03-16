package event

import (
	"event-management/hrm/storage"
	"fmt"
)


type EventStore interface {
	InsertEvent(s storage.Event) (*storage.Event, error)
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
	fmt.Println("=========cor=================")
	fmt.Println(s.PublishedAt)
	fmt.Println("==========================")
	event , err := ce.store.InsertEvent(s)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, fmt.Errorf("unable to insert event")
	}
	fmt.Println("==========================")
	fmt.Println(event.PublishedAt)
	fmt.Println("==========================")
	return event, nil
}