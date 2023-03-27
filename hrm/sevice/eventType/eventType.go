package eventType

import (
	"context"
	eventTypepb "event-management/gunk/v1/eventType"
	"event-management/hrm/storage"
	"sort"
)

type EventTypeCore interface {
	CreateEventType(s storage.EventTypes) (*storage.EventTypes, error)
	EditEventType(s storage.EventTypes) (*storage.EventTypes, error)
	UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error)
	DeleteEventType(s storage.EventTypes) error
	EventTypeListWithFilter(s storage.EventTypesFilter) ([]storage.EventTypes, error)
}

type EventTypeSvc struct {
	eventTypepb.UnimplementedEventTypeServiceServer
	Core EventTypeCore
}

func NewEventTypeSvc(ec EventTypeCore) *EventTypeSvc {
	return &EventTypeSvc{
		Core: ec,
	}
}

func (es EventTypeSvc) CreateEventType(ctx context.Context, r *eventTypepb.CreateEventTypeRequest) (*eventTypepb.CreateEventTypeResponse, error) {
	eventType := storage.EventTypes{
		EventTypeName: r.GetEventTypeName(),
	}

	if err := eventType.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.CreateEventType(eventType)
	if err != nil {
		return nil, err
	}

	return &eventTypepb.CreateEventTypeResponse{
		EventTypes: &eventTypepb.EventTypes{
			ID:        int32(et.ID),
			EventTypeName: eventType.EventTypeName,
		},
	}, nil
}

func (es EventTypeSvc) EditEventType(ctx context.Context, r *eventTypepb.EditEventTypeRequest) (*eventTypepb.EditEventTypeResponse, error) {
	eventTypeID := storage.EventTypes{
		ID: int(r.GetID()),
	}

	et, err := es.Core.EditEventType(eventTypeID)
	if err != nil {
		return nil, err
	}

	return &eventTypepb.EditEventTypeResponse{
		EventTypes: &eventTypepb.EventTypes{
			ID:        int32(et.ID),
			EventTypeName: et.EventTypeName,
		},
	}, nil
}

func (es EventTypeSvc) UpdateEventType(ctx context.Context, r *eventTypepb.UpdateEventTypeRequest) (*eventTypepb.UpdateEventTypeResponse, error) {
	eventType := storage.EventTypes{
		ID:        int(r.GetID()),
		EventTypeName: r.GetEventTypeName(),
	}

	if err := eventType.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	et, err := es.Core.UpdateEventType(eventType)
	if err != nil {
		return nil, err
	}

	return &eventTypepb.UpdateEventTypeResponse{
		EventTypes: &eventTypepb.EventTypes{
			ID:        int32(et.ID),
			EventTypeName: et.EventTypeName,
		},
	}, nil
}

func (es EventTypeSvc) DeleteEventType(ctx context.Context, r *eventTypepb.DeleteEventTypeRequest) (*eventTypepb.DeleteEventTypeResponse, error) {
	eventType := storage.EventTypes{
		ID: int(r.ID),
	}

	err := es.Core.DeleteEventType(eventType)
	if err != nil {
		return nil, err
	}

	return &eventTypepb.DeleteEventTypeResponse{}, nil
}

func (es EventTypeSvc) EventTypeList(ctx context.Context, r *eventTypepb.EventTypeListRequest) (*eventTypepb.EventTypeListResponse, error) {
	eventTypeList := storage.EventTypesFilter{
		SearchTerm: r.GetSearchTerm(),
	}

	et, err := es.Core.EventTypeListWithFilter(eventTypeList)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(et, func(i, j int) bool {
		return et[i].ID > et[j].ID
	})

	var allEventType []*eventTypepb.EventTypes
	for _, ae := range et {
		eventType := &eventTypepb.EventTypes{
			ID:        int32(ae.ID),
			EventTypeName: ae.EventTypeName,
		}
		allEventType = append(allEventType, eventType)
	}

	return &eventTypepb.EventTypeListResponse{
		EventTypeFilterList: &eventTypepb.EventTypeFilterList{
			AllEventType: allEventType,
			SearchTerm:   eventTypeList.SearchTerm,
		},
	}, nil
}
