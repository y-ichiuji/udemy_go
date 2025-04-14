package models

import (
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	Todos []Todo
}

type Session struct {
	ID int
	UUID string
	Email string
	UserID int
	CreatedAt time.Time
}

func (u *User) Create() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
		from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email, 
		&user.Password, 
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) Update() (err error) {
	cmd := `update users set 
		name = ?,
		email = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) Delete() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmdCreateSession := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmdCreateSession, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmdQuerySession := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmdQuerySession, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)

	return session, err
}

func (session *Session) Verify() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd, session.UUID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)

	if err != nil {
		valid = false
		return
	}

	if session.ID != 0 {
		valid = true
	}

	return valid, err
}

func (session *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, session.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (session *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at FROM users where id = ?`
	err = Db.QueryRow(cmd, session.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)
	return user, err
}