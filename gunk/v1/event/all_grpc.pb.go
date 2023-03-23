// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package eventpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	EditEvent(ctx context.Context, in *EditEventRequest, opts ...grpc.CallOption) (*EditEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
	EventList(ctx context.Context, in *EventListRequest, opts ...grpc.CallOption) (*EventListResponse, error)
	PublishedDateForm(ctx context.Context, in *PublishedDateFormRequest, opts ...grpc.CallOption) (*PublishedDateFormResponse, error)
	PublishedEvent(ctx context.Context, in *PublishedEventRequest, opts ...grpc.CallOption) (*PublishedEventResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) EditEvent(ctx context.Context, in *EditEventRequest, opts ...grpc.CallOption) (*EditEventResponse, error) {
	out := new(EditEventResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/EditEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) EventList(ctx context.Context, in *EventListRequest, opts ...grpc.CallOption) (*EventListResponse, error) {
	out := new(EventListResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/EventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) PublishedDateForm(ctx context.Context, in *PublishedDateFormRequest, opts ...grpc.CallOption) (*PublishedDateFormResponse, error) {
	out := new(PublishedDateFormResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/PublishedDateForm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) PublishedEvent(ctx context.Context, in *PublishedEventRequest, opts ...grpc.CallOption) (*PublishedEventResponse, error) {
	out := new(PublishedEventResponse)
	err := c.cc.Invoke(ctx, "/eventpb.EventService/PublishedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	EditEvent(context.Context, *EditEventRequest) (*EditEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
	EventList(context.Context, *EventListRequest) (*EventListResponse, error)
	PublishedDateForm(context.Context, *PublishedDateFormRequest) (*PublishedDateFormResponse, error)
	PublishedEvent(context.Context, *PublishedEventRequest) (*PublishedEventResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventServiceServer) EditEvent(context.Context, *EditEventRequest) (*EditEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEvent not implemented")
}
func (UnimplementedEventServiceServer) UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventServiceServer) DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventServiceServer) EventList(context.Context, *EventListRequest) (*EventListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventList not implemented")
}
func (UnimplementedEventServiceServer) PublishedDateForm(context.Context, *PublishedDateFormRequest) (*PublishedDateFormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishedDateForm not implemented")
}
func (UnimplementedEventServiceServer) PublishedEvent(context.Context, *PublishedEventRequest) (*PublishedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishedEvent not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_EditEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).EditEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/EditEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).EditEvent(ctx, req.(*EditEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_EventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).EventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/EventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).EventList(ctx, req.(*EventListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_PublishedDateForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishedDateFormRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).PublishedDateForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/PublishedDateForm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).PublishedDateForm(ctx, req.(*PublishedDateFormRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_PublishedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).PublishedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpb.EventService/PublishedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).PublishedEvent(ctx, req.(*PublishedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventpb.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventService_CreateEvent_Handler,
		},
		{
			MethodName: "EditEvent",
			Handler:    _EventService_EditEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventService_DeleteEvent_Handler,
		},
		{
			MethodName: "EventList",
			Handler:    _EventService_EventList_Handler,
		},
		{
			MethodName: "PublishedDateForm",
			Handler:    _EventService_PublishedDateForm_Handler,
		},
		{
			MethodName: "PublishedEvent",
			Handler:    _EventService_PublishedEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event-management/gunk/v1/event/all.proto",
}
