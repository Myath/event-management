package postgres

import (
	"event-management/hrm/storage"
	"fmt"
	"log"
)

const insertEventTypeQuery = `
INSERT INTO event_types(
	event_name
    )
VALUES(
	:event_name
    )RETURNING *;
`

func (p PostgresStorage) CreateEventType(s storage.EventTypes) (*storage.EventTypes, error){

	stmt, err := p.DB.PrepareNamed(insertEventTypeQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	if s.ID == 0 {
		return nil, fmt.Errorf("unable to insert eventType into db")
	}

	return &s, nil
}

const getEventTypeByIDQuery = `SELECT * FROM event_types WHERE id=$1 AND deleted_at IS NULL`

func (p PostgresStorage) GetEventTypeByID(id int) (*storage.EventTypes, error) {
	var s storage.EventTypes
	if err := p.DB.Get(&s, getEventTypeByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}

const updateEventTypeQuery = `UPDATE event_types
        SET	event_name = COALESCE(NULLIF(:event_name, ''), event_name),
		updated_at = CURRENT_TIMESTAMP
		WHERE id = :id
		RETURNING *;
	`

	func (p PostgresStorage) UpdateEventType(s storage.EventTypes) (*storage.EventTypes, error) {
		stmt, err := p.DB.PrepareNamed(updateEventTypeQuery)
		if err != nil {
			log.Fatalln(err)
		}
	
		if err := stmt.Get(&s, s); err != nil {
			return nil, err
		}
	
		if s.ID == 0 {
			return nil, fmt.Errorf("unable to insert eventType into db")
		}
	
		return &s, nil
	}