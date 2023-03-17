package postgres

import (
	"event-management/hrm/storage"
	"fmt"
	"log"
)


const insertEventQuery = `
INSERT INTO events(
	event_type_id,
	event_name,
	description,
	location,
	start_at,
	end_at,
	published_at
    )
VALUES(
	:event_type_id,
	:event_name,
	:description,
	:location,
	:start_at,
	:end_at,
	:published_at
    )RETURNING *;
`

func (p PostgresStorage) InsertEvent(s storage.Event) (*storage.Event, error){
	stmt, err := p.DB.PrepareNamed(insertEventQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	if s.ID == 0 {
		return nil, fmt.Errorf("unable to insert event into db")
	}

	return &s, nil
}

const getEventByIDQuery = `SELECT * FROM events WHERE id=$1 AND deleted_at IS NULL`

func (p PostgresStorage) GetEventByID(id int) (*storage.Event, error) {
	var s storage.Event
	if err := p.DB.Get(&s, getEventByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}

const updateEventQuery = `UPDATE events
        SET	event_type_id = :event_type_id,
		event_name = COALESCE(NULLIF(:event_name, ''), event_name),
		description = COALESCE(NULLIF(:description, ''), description),
		location = COALESCE(NULLIF(:location, ''), location),
		start_at = :start_at,
		end_at = :end_at,
		published_at = :published_at,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = :id
		RETURNING *;
	`

func (p PostgresStorage) UpdateEvent(s storage.Event) (*storage.Event, error) {
	stmt, err := p.DB.PrepareNamed(updateEventQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	if s.ID == 0 {
		return nil, fmt.Errorf("unable to update event into db")
	}

	return &s, nil
}