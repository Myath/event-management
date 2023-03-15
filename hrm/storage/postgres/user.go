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

func (ps PostgresStorage) GetUserByUsername(username string) (*storage.User, error) {
	var user storage.User
	if err := ps.DB.Get(&user, getUserByUsernameQuery, username); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to find user by username")
	}

	return &user, nil
}

const getUserByIDQuery = `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`

func (p PostgresStorage) GetUserByID(id int) (*storage.User, error) {
	var s storage.User
	if err := p.DB.Get(&s, getUserByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}

const updateUserQuery = `UPDATE users
        SET	first_name = COALESCE(NULLIF(:first_name, ''), first_name),
		last_name = COALESCE(NULLIF(:last_name, ''), last_name),
		username = COALESCE(NULLIF(:username, ''), username),
		email = COALESCE(NULLIF(:email, ''), email),
		password = COALESCE(NULLIF(:password, ''), password),
		is_active = :is_active
		WHERE id = :id
		RETURNING *;
	`

func (p PostgresStorage) UpdateUser(s storage.User) (*storage.User, error) {
	stmt, err := p.DB.PrepareNamed(updateUserQuery)
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

const deleteUserByIDQuery = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (p PostgresStorage) DeleteUserByID(id int) error {
	res, err := p.DB.Exec(deleteUserByIDQuery, id)
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
		return fmt.Errorf("unable to delete user")
	}

	return nil
}

const UserListQuery = `SELECT * FROM users WHERE deleted_at IS NULL AND (username ILIKE '%%' || $1 || '%%' or first_name ILIKE '%%' || $1 || '%%' or last_name ILIKE '%%' || $1 || '%%' or email ILIKE '%%' || $1 || '%%')`

func (p PostgresStorage) UserList(uf storage.UserFilter) ([]storage.User, error) {

	var user []storage.User
	if err := p.DB.Select(&user, UserListQuery, uf.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}