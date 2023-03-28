package handler

import (
	"log"
	"net/http"

	eventTypepb "event-management/gunk/v1/eventType"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)


type EventTypes struct {
	ID            int          `json:"id" form:"-" db:"id"`
	EventTypeName string       `json:"event_type_name" db:"event_type_name"`
}

type EventTypesFilter struct {
	AllEventsType []EventTypes
	SearchTerm string
	SeUserID string
}

func (e EventTypes) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.EventTypeName,
			validation.Required.Error("The event name is required."),
		),
	)
}

func (h Handler) EventTypeListForUser(w http.ResponseWriter, r *http.Request) {

	ad := r.FormValue("SearchTerm")


	u, err := h.hrmSvc.EventTypeList(r.Context(), &eventTypepb.EventTypeListRequest{
		SearchTerm: ad,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	allEventsType := []EventTypes{}
	if u != nil {
		for _, v := range u.GetEventTypeFilterList().AllEventType {

			e := EventTypes{
				ID:            int(v.GetID()),
				EventTypeName: v.GetEventTypeName(),
			}
			allEventsType = append(allEventsType, e)
		}
	}

	uID := h.SessionIDGet(w, r)

	Data := EventTypesFilter{
		AllEventsType: allEventsType,
		SearchTerm:    ad,
		SeUserID:      uID,
	}

	h.ParseEventTypeTemplate(w, Data)
}

func (h Handler) ParseEventTypeTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("eventtype.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}