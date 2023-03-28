package handler

import (
	"log"
	"net/http"
)


func(h Handler) Admin(w http.ResponseWriter, r *http.Request){
	h.pareseAdminDashboardTemplate(w, nil)
}

func (h Handler) pareseAdminDashboardTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("dashboard.html")
	if t == nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}