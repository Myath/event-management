package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"event-management/cms/handler"
	"event-management/utility"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/lib/pq"
)

//go:embed assets
var assetFiles embed.FS

//go:embed migrations
var migrationFiles embed.FS

var sessionManager *scs.SessionManager

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

	decoder := form.NewDecoder()

	lt := config.GetDuration("session.lifetime")
	it := config.GetDuration("session.idletime")
	sessionManager = scs.New()
	sessionManager.Lifetime = lt * time.Hour
	sessionManager.IdleTimeout = it * time.Minute
	sessionManager.Cookie.Name = "web-session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true

	postgreStorage, err := utility.NewPostgresStorage(config)
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

	var assetFS = fs.FS(assetFiles)
	staticFiles, err := fs.Sub(assetFS, "assets/src")
	if err != nil {
		log.Fatal(err)
	}

	templateFiles, err := fs.Sub(assetFS, "assets/templates")
	if err != nil {
		log.Fatal(err)
	}
	
	sessionManager.Store = utility.NewSQLXStore(postgreStorage.DB)

	p := config.GetInt("server.port")
	baseurl := config.GetString("server.baseurl")
	url := fmt.Sprintf("%d", p)

	hrmUrl := config.GetString("hrm.url")
	hrmConn, err := grpc.Dial(
		hrmUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(err)
	}

	chi := handler.NewHandler(sessionManager, decoder, hrmConn, baseurl, staticFiles, templateFiles)

	nosurfHandler := nosurf.New(chi)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", url))
	if err != nil {
		log.Fatalf("unable to listen port: %v", err)
	}

	fmt.Println("cms server running on: ", lis.Addr())
	if err := http.Serve(lis, nosurfHandler); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}
