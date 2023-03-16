package event

import (
	"context"
	eventpb "event-management/gunk/v1/event"
	"event-management/hrm/storage"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventCore interface {
	InsertEvent(s storage.Event) (*storage.Event, error)
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

func (es EventSvc) CreateEvent(ctx context.Context, r *eventpb.CreateEventRequest) (*eventpb.CreateEventResponse, error){
	fmt.Println("==========ssssssss===============")
	fmt.Println(r.GetPublishedAt().AsTime())
	fmt.Println(r.GetStartAt())
	fmt.Println("=========================")
	event := storage.Event{
		EventTypeId: int(r.GetEventTypeId()),
		EventName:   r.GetEventName(),
		Description: r.GetDescription(),
		Location:    r.GetLocation(),
		StartAt:     r.GetStartAt().AsTime(),
		EndAt:       r.GetEndAt().AsTime(),
		PublishedAt: r.GetPublishedAt().AsTime(),
	}

	if err := event.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.InsertEvent(event)
	if err != nil {
		return nil, err
	}
	fmt.Println("===========ssssssssssssss==============<<<<<<<<<<<<<<<<")
	fmt.Println(et.PublishedAt)
	fmt.Println("=========================")
	return &eventpb.CreateEventResponse{
		Event: &eventpb.Event{
			ID:          int32(et.ID),
			EventTypeId: int32(et.EventTypeId),
			EventName:   et.EventName,
			Description: et.Description,
			Location:    et.Location,
			StartAt:     timestamppb.New(et.StartAt),
			EndAt:       timestamppb.New(et.EndAt),
			PublishedAt: timestamppb.New(et.PublishedAt),
		},
	}, nil
}