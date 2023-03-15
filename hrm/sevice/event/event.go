package event

import (
	"context"
	eventpb "event-management/gunk/v1/event"
	"event-management/hrm/storage"
)

type EventCore interface {
	CreateEventType(s storage.EventTypes) (*storage.EventTypes, error)
	EditEventType(s storage.EventTypes) (*storage.EventTypes, error)
	UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error)
}

type EventSvc struct {
	eventpb.UnimplementedEventServiceServer
	Core EventCore
}

func NewEventSvc(ec EventCore) *EventSvc {
	return &EventSvc{
		Core: ec,
	}
}

func (es EventSvc) CreateEventType(ctx context.Context, r *eventpb.CreateEventTypeRequest) (*eventpb.CreateEventTypeResponse, error){
	event := storage.EventTypes{
		EventName: r.GetEventName(),
	}

	if err := event.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.CreateEventType(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.CreateEventTypeResponse{
		EventTypes: &eventpb.EventTypes{
			ID:        int32(et.ID),
			EventName: event.EventName,
		},
	}, nil
}

func (es EventSvc) EditEventType(ctx context.Context, r *eventpb.EditEventTypeRequest) (*eventpb.EditEventTypeResponse, error){
	eventTypeID := storage.EventTypes{
		ID: int(r.GetID()),
	}

	et, err := es.Core.EditEventType(eventTypeID)
	if err != nil {
		return nil, err
	}

	return &eventpb.EditEventTypeResponse{
		EventTypes: &eventpb.EventTypes{
			ID:        int32(et.ID),
			EventName: et.EventName,
		},
	}, nil
}

func (es EventSvc) UpdateEventType(ctx context.Context, r *eventpb.UpdateEventTypeRequest) (*eventpb.UpdateEventTypeResponse, error){
	eventType := storage.EventTypes{
		ID:        int(r.GetID()),
		EventName: r.GetEventName(),
	}

	if err := eventType.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.UpdateEventType(eventType)
	if err != nil {
		return nil, err
	}

	return &eventpb.UpdateEventTypeResponse{
		EventTypes: &eventpb.EventTypes{
			ID:        int32(et.ID),
			EventName: et.EventName,
		},
	}, nil
}

