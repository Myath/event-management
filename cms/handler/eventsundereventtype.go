package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	eventpb "event-management/gunk/v1/event"

	"github.com/go-chi/chi"
)

type EventUnderEventType struct {
	ID             int
	EventTypeId    int
	EventTypeName  string
	EventName      string
	Description    string
	Location       string
	StartAt        time.Time
	EndAt          time.Time
	PublishedAt    sql.NullTime
	PublishedAtStr string
	StartAtStr     string
	EndAtStr       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
}

type EventUnderEventTypeFilter struct {
	AllEventUnderEventType []EventUnderEventType
	SearchTerm             string
	SeUserID                 string
}

func (h Handler) EventUnderEventTypeForUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	eventID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	ad := r.FormValue("SearchTerm")

	u, err := h.hrmSvc.EventListUnderEvent(r.Context(), &eventpb.EventListUnderEventRequest{
		SearchTerm:  ad,
		EventTypeId: int32(eventID),
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	allEvents := []EventUnderEventType{}
	if u != nil {
		for _, v := range u.GetEvUnEvType().AllEvent {

			if !v.GetPublishedAt().AsTime().IsZero() {
				publishedAt := sql.NullTime{}
				if v.GetPublishedAt() != nil {
					publishedAt = sql.NullTime{
						Time:  v.GetPublishedAt().AsTime(),
						Valid: true,
					}
				}

				e := EventUnderEventType{
					ID:            int(v.ID),
					EventTypeId:   int(v.EventTypeId),
					EventTypeName: v.EventTypeName,
					EventName:     v.EventName,
					Description:   v.Description,
					Location:      v.Location,
					StartAt:       v.StartAt.AsTime(),
					StartAtStr:    v.StartAt.AsTime().Format("Jan 2, 2006 3:04PM"),
					EndAt:         v.EndAt.AsTime(),
					EndAtStr:      v.EndAt.AsTime().Format("Jan 2, 2006 3:04PM"),
					PublishedAt:   publishedAt,
				}
				if e.PublishedAt.Valid {
					e.PublishedAt.Time = e.PublishedAt.Time.Local() // convert to local timezone
					e.PublishedAtStr = e.PublishedAt.Time.Format("Jan 2, 2006 3:04PM")
				}

				allEvents = append(allEvents, e)
			}
		}
	}

	userID := h.SessionIDGet(w, r)

	Data := EventUnderEventTypeFilter{
		AllEventUnderEventType: allEvents,
		SearchTerm:             ad,
		SeUserID:                 userID,
	}

	h.ParseUnderEventTypeTemplate(w, Data)
}

func (h Handler) ParseUnderEventTypeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("eventsundereventtype.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
