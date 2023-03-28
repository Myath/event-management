package postgres

import (
	"event-management/hrm/storage"
	"fmt"
	"log"
)

const insertCommentQuery = `
INSERT INTO comments(
	user_id,
	event_id,
	comment
    )
VALUES(
	:user_id,
	:event_id,
	:comment
    )RETURNING *;
`

func (p PostgresStorage) InsertComment(s storage.Comments) (*storage.Comments, error) {
	stmt, err := p.DB.PrepareNamed(insertCommentQuery)
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

const getCommentIDQuery = `SELECT * FROM comments WHERE id=$1`

func (p PostgresStorage) GetCommentID(id int) (*storage.Comments, error) {
	var s storage.Comments
	if err := p.DB.Get(&s, getCommentIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}


const updateCommentQuery = `UPDATE comments
        SET	comment = COALESCE(NULLIF(:comment, ''), comment),
		updated_at = CURRENT_TIMESTAMP
		WHERE id = :id
		RETURNING *;
	`

func (p PostgresStorage) UpdateComment(s storage.Comments) (*storage.Comments, error) {
	stmt, err := p.DB.PrepareNamed(updateCommentQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	if s.ID == 0 {
		return nil, fmt.Errorf("unable to update comment into db")
	}

	return &s, nil
}

const CommentsOfEventQuery = `SELECT comments.event_id, comment, comments.id, users.username, comments.user_id, users.is_admin
FROM comments
INNER JOIN users ON comments.user_id = users.id
WHERE users.is_admin = false AND comments.event_id = $1
ORDER BY comments.created_at DESC
;`

func (p PostgresStorage) CommentsOfEvent(eventID int) ([]storage.Comments, error) {
	var comment []storage.Comments
	if err := p.DB.Select(&comment, CommentsOfEventQuery, eventID); err != nil {
		log.Println(err)
		return nil, err
	}
	return comment, nil
}