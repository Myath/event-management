package userevent

import (
	"context"
	usereventpb "event-management/gunk/v1/userEvent"
	"event-management/hrm/storage"
)

type UserEventCore interface {
	InsertUserEvent(s storage.UserEvent) (*storage.UserEvent, error)
	IsEventAvailable(int, int) (*storage.UserEvent, error)
	UpdateUserEvent(s storage.UserEvent) (*storage.UserEvent, error)
}

type UserEventSvc struct {
	usereventpb.UnimplementedUserEventServiceServer
	Core UserEventCore
}

func NewUserEventSvc(uec UserEventCore) *UserEventSvc {
	return &UserEventSvc{
		Core: uec,
	}
}

func (ues UserEventSvc) CreateUserEvent(ctx context.Context, r *usereventpb.CreateUserEventRequest) (*usereventpb.CreateUserEventResponse, error) {
	userEvent := storage.UserEvent{
		UserId:    int(r.GetUserId()),
		EventID:   int(r.GetEventID()),
		Status:    int(r.GetStatus()),
	}

	if err := userEvent.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	ue, err := ues.Core.InsertUserEvent(userEvent) 
	if err != nil {
		return nil, err
	}

	return &usereventpb.CreateUserEventResponse{
		UserEvent: &usereventpb.UserEvent{
			UserId:    int32(ue.UserId),
			EventID:   int32(ue.EventID),
			Status:    usereventpb.Status(ue.Status),
		},
	}, nil

}

func (ues UserEventSvc) IsEventAvailable(ctx context.Context, r *usereventpb.IsEventAvailableRequest) (*usereventpb.IsEventAvailableResponse, error){
	//TODO:: validation need


	ue, err := ues.Core.IsEventAvailable(int(r.GetUserId()), int(r.GetEventID())) 
	if err != nil {
		return nil, err
	}

	return &usereventpb.IsEventAvailableResponse{
		UserEvent: &usereventpb.UserEvent{
			UserId:    int32(ue.UserId),
			EventID:   int32(ue.EventID),
			Status:    usereventpb.Status(ue.Status),
		},
	}, nil

}

func (ues UserEventSvc) UpdateUserEvent(ctx context.Context, r *usereventpb.UpdateUserEventRequest) (*usereventpb.UpdateUserEventResponse, error) {
	userEvent := storage.UserEvent{
		UserId:    int(r.UserEvent.GetUserId()),
		EventID:   int(r.UserEvent.GetEventID()),
		Status:    int(r.UserEvent.GetStatus()),
	}

	if err := userEvent.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	ue, err := ues.Core.UpdateUserEvent(userEvent) 
	if err != nil {
		return nil, err
	}

	return &usereventpb.UpdateUserEventResponse{
		UserEvent: &usereventpb.UserEvent{
			UserId:    int32(ue.UserId),
			EventID:   int32(ue.EventID),
			Status:    usereventpb.Status(ue.Status),
		},
	}, nil

}
