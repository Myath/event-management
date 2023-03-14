package postgres

import (
	"event-management/hrm/storage"
	"fmt"
	"log"
)
const insertRegisterQuery = `
	INSERT INTO users(
		first_name,
		last_name,
		username,
		email,
		password
		)  
	VALUES(
		:first_name,
		:last_name,
		:username,
		:email,
		:password
		)RETURNING *;
	`

func (p PostgresStorage) Register(s storage.User) (*storage.User, error) {

	stmt, err := p.DB.PrepareNamed(insertRegisterQuery)
	if err != nil {
		log.Fatalln(err)
	}


	if err := stmt.Get(&s, s); err != nil {
		return nil, err
	}

	if s.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}

	return &s, nil
}

const getUserByUsernameQuery = `SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL`

func (ps PostgresStorage) GetUserByUsername(usernanme string) (*storage.User, error) {
	var user storage.User
	if err := ps.DB.Get(&user, getUserByUsernameQuery, usernanme); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to find user by username")
	}

	return &user, nil
}