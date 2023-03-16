package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"strings"

	eventpb "event-management/gunk/v1/event"
	eventTypepb "event-management/gunk/v1/eventType"
	userpb "event-management/gunk/v1/user"
	et "event-management/hrm/core/eventType"
	ev "event-management/hrm/core/event"
	cu "event-management/hrm/core/user"
	"event-management/hrm/sevice/eventType"
	"event-management/hrm/sevice/event"
	"event-management/hrm/sevice/user"
	"event-management/hrm/storage/postgres"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:embed migrations
var migrationFiles embed.FS

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	port := config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen port: %v", err)
	}

	postgreStorage, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}

	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}

	if err := goose.Up(postgreStorage.DB.DB, "migrations"); err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	userCore := cu.NewCoreUser(postgreStorage)
	userSvc := user.NewUserSvc(userCore)
	userpb.RegisterUserServiceServer(grpcServer, userSvc)

	eventTypeCore := et.NewCoreEventType(postgreStorage)
	eventTypeSvc := eventType.NewEventTypeSvc(eventTypeCore)
	eventTypepb.RegisterEventTypeServiceServer(grpcServer,eventTypeSvc)

	eventCore := ev.NewCoreEvent(postgreStorage)
	eventSvc := event.NewEventSvc(eventCore)
	eventpb.RegisterEventServiceServer(grpcServer,eventSvc)

	// start reflection server
	reflection.Register(grpcServer)

	fmt.Println("usermgm server running on: ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}