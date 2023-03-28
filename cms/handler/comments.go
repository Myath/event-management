package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	commentpb "event-management/gunk/v1/comment"
	eventpb "event-management/gunk/v1/event"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

type Events struct {
	ID             int
	EventName      string
	Description    string
	Location       string
	StartAt        time.Time
	EndAt          time.Time
	PublishedAt    sql.NullTime
	PublishedAtStr string
	StartAtStr     string
	EndAtStr       string
}

type EventsComments struct {
	ID        int
	UserId    int
	StrUserId string
	Username  string
	IsAdmin   bool
	EventID   int
	Comment   string
	FormError map[string]error
}

type UserEventStatus struct {
	UserId  int
	EventID int
	Status  string
}

func (c EventsComments) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.UserId,
			validation.Required.Error("The userId is required."),
		),
		validation.Field(&c.EventID,
			validation.Required.Error("The eventID name field is required."),
		),
		validation.Field(&c.Comment,
			validation.Required.When(c.ID == 0).Error("The Comment field is required."),
		),
	)
}

type EventsCommentStatusForm struct {
	SingleEvent Events
	Comments    []EventsComments
	UserEvent   UserEventStatus
	SeUserID    string
	FormError   map[string]error
	CSRFToken   string
}

func (h Handler) EventsCommentStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	eventID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	uID := h.SessionIDGet(w, r)

	u, err := h.hrmSvc.EditEvent(r.Context(), &eventpb.EditEventRequest{
		ID: int32(eventID),
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var event Events
	if u != nil {
		if !u.Event.GetPublishedAt().AsTime().IsZero() {
			publishedAt := sql.NullTime{}
			if u.Event.GetPublishedAt() != nil {
				publishedAt = sql.NullTime{
					Time:  u.Event.GetPublishedAt().AsTime(),
					Valid: true,
				}
			}

			event = Events{
				ID:          int(u.Event.ID),
				EventName:   u.Event.EventName,
				Description: u.Event.Description,
				Location:    u.Event.Location,
				StartAt:     u.Event.StartAt.AsTime(),
				StartAtStr:  u.Event.StartAt.AsTime().Format("Jan 2, 2006 3:04PM"),
				EndAtStr:    u.Event.EndAt.AsTime().Format("Jan 2, 2006 3:04PM"),
				EndAt:       u.Event.EndAt.AsTime(),
				PublishedAt: publishedAt,
			}
			if event.PublishedAt.Valid {
				event.PublishedAt.Time = event.PublishedAt.Time.Local() // convert to local timezone
				event.PublishedAtStr = event.PublishedAt.Time.Format("Jan 2, 2006 3:04PM")
			}
		}
	}

	ct, err := h.hrmSvc.CommentListOfEVents(r.Context(), &commentpb.CommentListOfEVentsRequest{
		EventID: int32(eventID),
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	eventComments := []EventsComments{}
	if ct != nil {
		for _, v := range ct.GetCmtOfEv() {

			c := EventsComments{
				ID:        int(v.ID),
				UserId:    int(v.UserId),
				StrUserId: "",
				Username:  v.Username,
				IsAdmin:   v.IsAdmin,
				EventID:   int(v.EventID),
				Comment:   v.Comment,
				FormError: map[string]error{},
			}
			c.StrUserId = strconv.Itoa(c.UserId)

			eventComments = append(eventComments, c)
		}
	}

	Data := EventsCommentStatusForm{
		SingleEvent: event,
		Comments:    eventComments,
		// UserEvent: userEvent,
		SeUserID:  uID,
		CSRFToken:  nosurf.Token(r),
	}

	h.ParseCommentsTemplate(w, Data)
}

func (h Handler) ParseCommentsTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("commentdetails.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h Handler) AddCommentByUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	eventID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	uID := h.SessionIDGet(w, r)
	userID, err := strconv.Atoi(uID)
	if err != nil {
		log.Fatal(err)
	}

	comment := EventsComments{}
	if err := h.decoder.Decode(&comment, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	_, erR := h.hrmSvc.CreateComment(r.Context(), &commentpb.CreateCommentRequest{
		UserId:  int32(userID),
		EventID: int32(eventID),
		Comment: comment.Comment,
	})
	if erR != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s/%d", "/commentdetails", eventID), http.StatusSeeOther)
}
