package handler

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	eventpb "event-management/gunk/v1/event"
)

type Event struct {
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

type EventFilter struct {
	AllEvents  []Event
	SearchTerm string
	SeUserID     string
}

func (h Handler) EventForUser(w http.ResponseWriter, r *http.Request) {

	ad := r.FormValue("SearchTerm")

	u, err := h.hrmSvc.EventList(r.Context(), &eventpb.EventListRequest{
		SearchTerm: ad,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	allEvents := []Event{}
	if u != nil {
		for _, v := range u.GetEventFilterList().AllEvent {

			if !v.GetPublishedAt().AsTime().IsZero() {
				publishedAt := sql.NullTime{}
				if v.GetPublishedAt() != nil {
					publishedAt = sql.NullTime{
						Time:  v.GetPublishedAt().AsTime(),
						Valid: true,
					}
				}

				e := Event{
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

	Data := EventFilter{
		AllEvents:  allEvents,
		SearchTerm: ad,
		SeUserID:     userID,
	}

	h.ParseEventTemplate(w, Data)
}

func (h Handler) ParseEventTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("events-user.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
