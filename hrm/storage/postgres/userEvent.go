package postgres

import (
	"database/sql"
	"event-management/hrm/storage"
	// "fmt"
	"log"
)

const insertUserEventQuery = `
INSERT INTO user_events (
	user_id,
	event_id,
	status
    )
VALUES(
	:user_id,
	:event_id,
	:status
) RETURNING *;
`

func (p PostgresStorage) InsertUserEvent(s storage.UserEvent) (*storage.UserEvent, error) {
	stmt, err := p.DB.PrepareNamed(insertUserEventQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	return &s, nil
}

const getUserEventByIDQuery = `SELECT * FROM user_events WHERE user_id=$1 AND event_id = $2`

func (p PostgresStorage) GetUserEventByIDQuery(userId, eventId int) (*storage.UserEvent, error) {
	var s storage.UserEvent
	if err := p.DB.Get(&s, getUserEventByIDQuery, userId, eventId); err != nil || err == sql.ErrNoRows{
		log.Println(err)
		return nil, err
	}

	return &s, nil
}


const updateUserEventQuery = `UPDATE user_events
        SET	user_id = :user_id,
		event_id = :event_id,
		status = :status,
		updated_at = CURRENT_TIMESTAMP
		WHERE user_id= :user_id AND event_id = :event_id
		RETURNING *;
	`

func (p PostgresStorage) UpdateUserEvent(s storage.UserEvent) (*storage.UserEvent, error) {
	stmt, err := p.DB.PrepareNamed(updateUserEventQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	return &s, nil
}