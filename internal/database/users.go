package database

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id                 int       `json:"id"`
	Email              string    `json:"email"`
	Name               string    `json:"name"`
	Password           string    `json:"-"`
	Role               string    `json:"role"`
	Verified           bool      `json:"verified"`
	VerifyToken        string    `json:"verify_token"`
	VerifyTokenExpires time.Time `json:"verify_token_expires"`
}

func generateVerifyToken() string {
	bytes := make([]byte, 12)
	_, err := rand.Read(bytes)

	if err != nil {
		fmt.Println("Failed to generate random token:", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

// Insert adds a new user to the database.
// @Summary Add a new user
// @Description Insert a new user into the database with verification details.
// @Tags User
// @Param user body User true "User data"
// @Success 200 {object} User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 預設產生未驗證狀態
	user.Verified = false
	user.VerifyToken = generateVerifyToken()                 // 生成隨機驗證 token
	user.VerifyTokenExpires = time.Now().Add(24 * time.Hour) // 設定驗證期限

	query := `
		INSERT INTO users (email, password, name, role, verified, verify_token, verify_token_expires)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	return m.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name, user.Role, user.Verified, user.VerifyToken, user.VerifyTokenExpires).Scan(&user.Id)
}

func (m *UserModel) getUser(query string, args ...interface{}) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Email, &user.Name, &user.Password, &user.Role, &user.Verified, &user.VerifyToken, &user.VerifyTokenExpires,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Get retrieves a user by ID.
// @Summary Get user by ID
// @Description Retrieve a user from the database using their ID.
// @Tags User
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [get]
func (m *UserModel) Get(id int) (*User, error) {
	query := `
		SELECT id, email, name, password, role, verified, verify_token, verify_token_expires
		FROM users
		WHERE id = $1
	`
	return m.getUser(query, id)
}

// GetByEmail retrieves a user by email.
// @Summary Get user by email
// @Description Retrieve a user from the database using their email.
// @Tags User
// @Param email query string true "User email"
// @Success 200 {object} User
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/email [get]
func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, email, name, password, role, verified, verify_token, verify_token_expires
		FROM users
		WHERE email = $1
	`
	return m.getUser(query, email)
}

// Update updates the user's name and password. Email cannot be updated.
// @Summary Update user details
// @Description Update the name and password of a user. Email cannot be updated.
// @Tags User
// @Param id path int true "User ID"
// @Param name body string false "New name"
// @Param password body string false "New password"
// @Success 200 {object} User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [put]
func (m *UserModel) Update(id int, name, password string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Ensure email is existing
	var email string
	err := m.DB.QueryRowContext(ctx, "SELECT email FROM users WHERE id = $1", id).Scan(&email)
	if err != nil {
		return nil, err
	}

	query := `
		UPDATE users
		SET name = COALESCE(NULLIF($1, ''), name),
		    password = COALESCE(NULLIF($2, ''), password)
		WHERE id = $3
		RETURNING id, email, name, password, role, verified, verify_token, verify_token_expires
	`

	var user User
	err = m.DB.QueryRowContext(ctx, query, name, password, id).Scan(
		&user.Id, &email, &user.Name, &user.Password, &user.Role, &user.Verified, &user.VerifyToken, &user.VerifyTokenExpires,
	)

	if err != nil {
		return nil, err
	}

	user.Email = email
	return &user, nil
}
