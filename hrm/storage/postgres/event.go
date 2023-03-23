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

const deleteEventByIDQuery = `UPDATE events SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (p PostgresStorage) DeleteEventByID(id int) error {
	res, err := p.DB.Exec(deleteEventByIDQuery, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return fmt.Errorf("unable to delete event")
	}

	return nil
}

const listEventQuery = `SELECT events.id, events.event_type_id, event_types.event_type_name, event_name, description, location, start_at, end_at, published_at
FROM events
INNER JOIN event_types ON events.event_type_id = event_types.id
WHERE events.deleted_at IS NULL AND event_types.deleted_at IS NULL AND (event_type_name ILIKE '%%' || $1 || '%%' or event_name ILIKE '%%' || $1 || '%%' or description ILIKE '%%' || $1 || '%%' or location ILIKE '%%' || $1 || '%%')
ORDER BY events.created_at DESC
;`

func (p PostgresStorage) ListEvent(uf storage.EventFilter) ([]storage.Event, error) {

	var event []storage.Event
	if err := p.DB.Select(&event, listEventQuery, uf.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return event, nil
}

const getPublishedDateByIDQuery = `SELECT id, published_at FROM events WHERE id=$1 AND deleted_at IS NULL`

func (p PostgresStorage) GetPublishedDateByID(id int) (*storage.Event, error) {
	var s storage.Event
	if err := p.DB.Get(&s, getPublishedDateByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}

const publishedEventQuery = `UPDATE events
        SET	published_at = :published_at,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = :id
		RETURNING *;
	`

func (p PostgresStorage) PublishedEvent(s storage.Event) (*storage.Event, error) {
	stmt, err := p.DB.PrepareNamed(publishedEventQuery)
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