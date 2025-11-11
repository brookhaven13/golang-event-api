package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (email, password, name, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	return m.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name, user.Role).Scan(&user.Id)
}

func (m *UserModel) getUser(query string, args ...interface{}) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	query := `
		SELECT id, email, name, password, role
		FROM users
		WHERE id = $1
	`
	return m.getUser(query, id)
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT * FROM users
		WHERE email = $1
	`
	return m.getUser(query, email)
}

func (m *UserModel) Update(id int, email, name, password string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE users
		SET email = COALESCE(NULLIF($1, ''), email),
		    name = COALESCE(NULLIF($2, ''), name),
		    password = COALESCE(NULLIF($3, ''), password)
		WHERE id = $4
		RETURNING id, email, name, password, role
	`

	var user User
	err := m.DB.QueryRowContext(ctx, query, email, name, password, id).Scan(
		&user.Id, &user.Email, &user.Name, &user.Password, &user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
