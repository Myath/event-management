package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	userpb "event-management/gunk/v1/user"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/form"
	"google.golang.org/grpc"
)

type hrmService struct{
	userpb.UserServiceClient
}

type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	hrmSvc hrmService
	Templates      *template.Template
	baseurl        string
}

// const (
// 	adminLoginPath = "/adminLogin"
// )

func NewHandler(sm *scs.SessionManager, formdecoder *form.Decoder, hrmConn *grpc.ClientConn, baseurl string) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formdecoder,
		hrmSvc:         hrmService{userpb.NewUserServiceClient(hrmConn)},
		baseurl:        baseurl,
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

	//Make Template prefix variable
	// assetsPrefixForSubjectEdit := "/subject/edit/static/"
	// assetsPrefixForSubjectUpdate := "/subject/update/static/"
	// assetsPrefixForStudentEdit := "/admitstudent/edit/static/"
	// assetsPrefixForStudentUpdate := "/admitstudent/update/static/"
	// assetsPrefixForClassEdit := "/class/edit/static/"
	// assetsPrefixForClassUpdate := "/class/update/static/"
	// assetsPrefixForAdminEdit := "/admin/edit/static/"
	// assetsPrefixForAdminUpdate := "/admin/update/static/"
	// assetsPrefixForAddMark := "/addmark/static/"
	// assetsPrefixForProfile := "/profile/static/"
	// assetsPrefixForEditMark:= "/editmark/static/"
	// assetsPrefixForUpdateMark:= "/updatemark/static/"

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(VerbMethod)

	// r.Get("/admincreate", h.AdminCreate)
	// r.Post("/admincreate", h.AdminCreateProcess)

	// r.Group(func(r chi.Router) {
	// 	r.Use(sm.LoadAndSave)
	// 	r.Use(h.AuthenticationForLogin)
	// 	r.Get(adminLoginPath, h.AdminLogin)
	// 	r.Post(adminLoginPath, h.AdminLoginProcess)
	// })

	// For Template Asset Prefixes
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "assets/src"))
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForSubjectEdit+"*", http.StripPrefix(assetsPrefixForSubjectEdit, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForSubjectUpdate+"*", http.StripPrefix(assetsPrefixForSubjectUpdate, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForStudentEdit+"*", http.StripPrefix(assetsPrefixForStudentEdit, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForStudentUpdate+"*", http.StripPrefix(assetsPrefixForStudentUpdate, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForClassEdit+"*", http.StripPrefix(assetsPrefixForClassEdit, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForClassUpdate+"*", http.StripPrefix(assetsPrefixForClassUpdate, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForAdminEdit+"*", http.StripPrefix(assetsPrefixForAdminEdit, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForAdminUpdate+"*", http.StripPrefix(assetsPrefixForAdminUpdate, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForAddMark+"*", http.StripPrefix(assetsPrefixForAddMark, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForProfile+"*", http.StripPrefix(assetsPrefixForProfile, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForEditMark+"*", http.StripPrefix(assetsPrefixForEditMark, http.FileServer(filesDir)))
	// r.Handle(assetsPrefixForUpdateMark+"*", http.StripPrefix(assetsPrefixForUpdateMark, http.FileServer(filesDir)))

	// r.Group(func(r chi.Router) {
	// 	r.Use(sm.LoadAndSave)
	// 	r.Use(h.Authentication)
	// 	r.Get("/admin/edit/{id:[0-9]+}", h.AdminEdit)
	// 	r.Put("/admin/update/{id:[0-9]+}", h.AdminUpdate)
	// 	r.Get("/admin/delete/{id:[0-9]+}", h.DeleteAdmin)
	// })

	// r.Group(func(r chi.Router) {
	// 	r.Use(sm.LoadAndSave)
	// 	r.Use(h.Authentication)
	// 	r.Get("/dashboard", h.Dashboard)
	// 	r.Get("/adminlist", h.AdminList)

	// 	// Class Route
	// 	r.Get("/classcreate", h.ClassCreate)
	// 	r.Post("/classcreate", h.ClassCreateProcess)
	// 	r.Get("/classlist", h.ClassList)
	// 	r.Get("/class/edit/{id:[0-9]+}", h.EditClass)
	// 	r.Put("/class/update/{id:[0-9]+}", h.ClassUpdate)
	// 	r.Get("/class/delete/{id:[0-9]+}", h.DeleteClass)

	// 	// Subject Route
	// 	r.Get("/subjectcreate", h.SubjectCreate)
	// 	r.Post("/subjectcreate", h.SubjectCreateProcess)
	// 	r.Get("/subjectlist", h.SubjectList)
	// 	r.Get("/subject/edit/{id:[0-9]+}", h.SubjectEdit)
	// 	r.Put("/subject/update/{id:[0-9]+}", h.SubjectUpdate)
	// 	r.Get("/subject/delete/{id:[0-9]+}", h.DeleteSubject)

	// 	// Admit Student
	// 	r.Get("/admitstudent", h.AdmitStudent)
	// 	r.Post("/admitstudent", h.AdmitStudentProcess)
	// 	r.Get("/admitstudentlist", h.AdmitStudentlist)
	// 	r.Get("/admitstudent/edit/{id:[0-9]+}", h.AdmitStudentEdit)
	// 	r.Put("/admitstudent/update/{id:[0-9]+}", h.AdmitStudentUpdate)
	// 	r.Get("/admitstudent/delete/{id:[0-9]+}", h.DeleteAdmitStudent)

	// 	r.Get("/addmark/{id:[0-9]+}", h.AddMark)
	// 	r.Post("/markstore", h.Markstore)
	// 	r.Get("/profile/{id:[0-9]+}", h.Profile)
	// 	r.Get("/editmark/{id:[0-9]+}", h.EditMark)
	// 	r.Post("/updatemark/{id:[0-9]+}", h.MarkUpdate)
	// 	r.Get("/deletemark/{id:[0-9]+}", h.DeleteMark)

	// 	r.Get("/addadmin", h.AdminAdd)
	// 	r.Post("/addadmin", h.AddAdminCreateProcess)
	// 	r.Get("/adminlogout", h.AdminLogOut)
	// })

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
		username := h.sessionManager.GetString(r.Context(), "username")
		if username == "" {
			http.Redirect(w, r, "/adminLogin", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h Handler) AuthenticationForLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := h.sessionManager.GetString(r.Context(), "username")
		if username != "" {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
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

	newFS := os.DirFS("assets/templates")
	tmpl := template.Must(templates.ParseFS(newFS, "*/*.html", "*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
