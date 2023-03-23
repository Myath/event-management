package event

import (
	"context"
	"database/sql"

	// "database/sql"
	eventpb "event-management/gunk/v1/event"
	"event-management/hrm/storage"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventCore interface {
	InsertEvent(s storage.Event) (*storage.Event, error)
	EditUser(s storage.Event) (*storage.Event, error)
	UpdateEvent(s storage.Event) (*storage.Event, error)
	DeleteEvent(s storage.Event) error
	EventListWithFilter(s storage.EventFilter) ([]storage.Event, error)
	PublishedDateForm(s storage.Event) (*storage.Event, error)
	PublishedEvent(s storage.Event) (*storage.Event, error)
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

func (es EventSvc) CreateEvent(ctx context.Context, r *eventpb.CreateEventRequest) (*eventpb.CreateEventResponse, error) {

	publishedAt := sql.NullTime{}
	if r.GetPublishedAt() != nil {
		publishedAt = sql.NullTime{
			Time:  r.GetPublishedAt().AsTime(),
			Valid: true,
		}
	}

	event := storage.Event{
		EventTypeId: int(r.GetEventTypeId()),
		EventName:   r.GetEventName(),
		Description: r.GetDescription(),
		Location:    r.GetLocation(),
		StartAt:     r.GetStartAt().AsTime(),
		EndAt:       r.GetEndAt().AsTime(),
		PublishedAt: publishedAt,
	}

	if err := event.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.InsertEvent(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.CreateEventResponse{
		Event: &eventpb.Event{
			ID:          int32(et.ID),
			EventTypeId: int32(et.EventTypeId),
			EventName:   et.EventName,
			Description: et.Description,
			Location:    et.Location,
			StartAt:     timestamppb.New(et.StartAt),
			EndAt:       timestamppb.New(et.EndAt),
			PublishedAt: timestamppb.New(et.PublishedAt.Time),
		},
	}, nil
}

func (es EventSvc) EditEvent(ctx context.Context, r *eventpb.EditEventRequest) (*eventpb.EditEventResponse, error) {
	eventID := storage.Event{
		ID: int(r.GetID()),
	}

	et, err := es.Core.EditUser(eventID)
	if err != nil {
		return nil, err
	}

	return &eventpb.EditEventResponse{
		Event: &eventpb.Event{
			ID:          int32(et.ID),
			EventTypeId: int32(et.EventTypeId),
			EventName:   et.EventName,
			Description: et.Description,
			Location:    et.Location,
			StartAt:     timestamppb.New(et.StartAt),
			EndAt:       timestamppb.New(et.EndAt),
			PublishedAt: timestamppb.New(et.PublishedAt.Time),
		},
	}, nil
}

func (es EventSvc) UpdateEvent(ctx context.Context, r *eventpb.UpdateEventRequest) (*eventpb.UpdateEventResponse, error) {
	publishedAt := sql.NullTime{}
	if r.Event.GetPublishedAt() != nil {
		publishedAt = sql.NullTime{
			Time:  r.Event.GetPublishedAt().AsTime(),
			Valid: true,
		}
	}

	event := storage.Event{
		ID:          int(r.Event.GetID()),
		EventTypeId: int(r.Event.GetEventTypeId()),
		EventName:   r.Event.GetEventName(),
		Description: r.Event.GetDescription(),
		Location:    r.Event.GetLocation(),
		StartAt:     r.Event.StartAt.AsTime(),
		EndAt:       r.Event.EndAt.AsTime(),
		PublishedAt: publishedAt,
	}

	if err := event.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.UpdateEvent(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.UpdateEventResponse{
		Event: &eventpb.Event{
			ID:          int32(et.ID),
			EventTypeId: int32(et.EventTypeId),
			EventName:   et.EventName,
			Description: et.Description,
			Location:    et.Location,
			StartAt:     timestamppb.New(et.StartAt),
			EndAt:       timestamppb.New(et.EndAt),
			PublishedAt: timestamppb.New(et.PublishedAt.Time),
		},
	}, nil
}

func (es EventSvc) DeleteEvent(ctx context.Context, r *eventpb.DeleteEventRequest) (*eventpb.DeleteEventResponse, error) {
	event := storage.Event{
		ID: int(r.ID),
	}

	err := es.Core.DeleteEvent(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.DeleteEventResponse{}, nil
}


func (es EventSvc) EventList(ctx context.Context, r *eventpb.EventListRequest) (*eventpb.EventListResponse, error){
	eventList := storage.EventFilter{
		SearchTerm: r.GetSearchTerm(),
	}

	ev, err :=  es.Core.EventListWithFilter(eventList)
	if err != nil {
		return nil, err
	}

	var allEvents []*eventpb.EventList
	for _, evt := range ev {
		event := &eventpb.EventList{
			ID:          int32(evt.ID),
			EventTypeId: int32(evt.EventTypeId),
			EventTypeName : evt.EventTypeName,
			EventName:   evt.EventName,
			Description: evt.Description,
			Location:    evt.Location,
			StartAt:     timestamppb.New(evt.StartAt),
			EndAt:       timestamppb.New(evt.EndAt),
			PublishedAt: timestamppb.New(evt.PublishedAt.Time),
		}
		allEvents = append(allEvents, event)
	}
	return &eventpb.EventListResponse{
		EventFilterList: &eventpb.EventFilterList{
			AllEvent:   allEvents,
			SearchTerm: eventList.SearchTerm,
		},
	}, nil
}

func (es EventSvc) PublishedDateForm(ctx context.Context, r *eventpb.PublishedDateFormRequest) (*eventpb.PublishedDateFormResponse, error){
	eventID := storage.Event{
		ID: int(r.GetID()),
	}

	et, err := es.Core.PublishedDateForm(eventID)
	if err != nil {
		return nil, err
	}

	return &eventpb.PublishedDateFormResponse{
		ID:          int32(et.ID),
		PublishedAt: timestamppb.New(et.PublishedAt.Time),
	}, nil
}

func (es EventSvc) PublishedEvent(ctx context.Context, r *eventpb.PublishedEventRequest) (*eventpb.PublishedEventResponse, error){
	publishedAt := sql.NullTime{}
	if r.GetPublishedAt() != nil {
		publishedAt = sql.NullTime{
			Time:  r.GetPublishedAt().AsTime(),
			Valid: true,
		}
	}

	event := storage.Event{
		ID:          int(r.GetID()),
		PublishedAt: publishedAt,
	}

	et, err := es.Core.PublishedEvent(event)
	if err != nil {
		return nil, err
	}

	return &eventpb.PublishedEventResponse{
		ID:          int32(et.ID),
		PublishedAt: timestamppb.New(et.PublishedAt.Time),
	}, nil
}