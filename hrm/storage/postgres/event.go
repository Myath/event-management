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
	fmt.Println("==========================")
	fmt.Println(s.PublishedAt)
	fmt.Println("==========================")

	return &s, nil
}