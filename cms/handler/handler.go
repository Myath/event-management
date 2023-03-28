package handler

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	eventpb "event-management/gunk/v1/event"
	eventTypepb "event-management/gunk/v1/eventType"
	userpb "event-management/gunk/v1/user"
	commentpb "event-management/gunk/v1/comment"
	usereventpb "event-management/gunk/v1/userEvent"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/form"
	"google.golang.org/grpc"
)

type hrmService struct {
	userpb.UserServiceClient
	eventpb.EventServiceClient
	eventTypepb.EventTypeServiceClient
	commentpb.CommentServiceClient
	usereventpb.UserEventServiceClient
}

type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	hrmSvc         hrmService
	Templates      *template.Template
	baseurl        string
	staticFiles    fs.FS
	templateFiles  fs.FS
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

type ErrorPage struct {
	Code    int
	Message string
}

func (h Handler) Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	ep := ErrorPage{
		Code:    code,
		Message: error,
	}

	tf := "default"
	switch code {
	case 400, 401, 402, 403, 404:
		tf = "4xx"
	case 500, 501, 503:
		tf = "5xx"
	}

	tpl := fmt.Sprintf("%s.html", tf)
	t := h.Templates.Lookup(tpl)
	if t == nil {
		log.Fatalln("unable to find template")
	}

	if err := t.Execute(w, ep); err != nil {
		log.Fatalln(err)
	}
}

func NewHandler(sm *scs.SessionManager, formdecoder *form.Decoder, hrmConn *grpc.ClientConn, baseurl string, staticFiles, templateFiles fs.FS) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formdecoder,
		hrmSvc: hrmService{userpb.NewUserServiceClient(hrmConn),
			eventpb.NewEventServiceClient(hrmConn),
			eventTypepb.NewEventTypeServiceClient(hrmConn),
			commentpb.NewCommentServiceClient(hrmConn),
			usereventpb.NewUserEventServiceClient(hrmConn),
		},
		baseurl:       baseurl,
		staticFiles:   staticFiles,
		templateFiles: templateFiles,
	}

	h.ParseTemplates()
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Headers"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(VerbMethod)

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(h.staticFiles))))
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/", h.Home)
		r.Get("/register", h.Register)
		r.Post("/register", h.RegisterPost)
		r.Get("/login", h.Login)
		r.Post("/login", h.LoginPostHandler)
		r.Get("/logout", h.Logout)

	})

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/dashboard", h.Admin)
	})

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/eventtype", h.EventTypeListForUser)
		r.Get("/events", h.EventForUser)
		r.Get("/evunderevtype/{id:[0-9]+}", h.EventUnderEventTypeForUser)
		r.Get("/commentdetails/{id:[0-9]+}", h.EventsCommentStatus)
		r.Post("/commentadd/{id:[0-9]+}", h.AddCommentByUser)

	})

	return r
}

func VerbMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch strings.ToLower(r.PostFormValue("_method")) {
			case "put":
				r.Method = http.MethodPut
			case "patch":
				r.Method = http.MethodPatch
			case "delete":
				r.Method = http.MethodDelete
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (h Handler) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := h.sessionManager.GetString(r.Context(), "userID")
		uID, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if uID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ParseTemplates() error {
	templates := template.New("web-templates").Funcs(template.FuncMap{
		"assetGenerate": func(n string) string {
			return fmt.Sprintf("%s/%s", h.baseurl, strings.Trim(n, "/"))
		},
	}).Funcs(sprig.FuncMap())

	tmpl := template.Must(templates.ParseFS(h.templateFiles, "*/*.html", "*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
