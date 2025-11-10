package database

import (
	"context"
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int       `json:"id"`
	OwnerId     int       `json:"-"`
	Owner       *User     `json:"owner,omitempty"`
	Name        string    `json:"name" binding:"required,min=3"`
	Description string    `json:"description" binding:"required,min=10"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"required,min=3"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `
		INSERT INTO events (owner_id, name, description, date, location)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := m.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `
		SELECT e.id, e.owner_id, e.name, e.description, e.date, e.location,
		       u.id, u.email, u.name, u.role
		FROM events e
		LEFT JOIN users u ON e.owner_id = u.id
	`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event
		var owner User

		err := rows.Scan(
			&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location,
			&owner.Id, &owner.Email, &owner.Name, &owner.Role,
		)
		if err != nil {
			return nil, err
		}

		event.Owner = &owner
		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (m *EventModel) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT e.id, e.owner_id, e.name, e.description, e.date, e.location,
		       u.id, u.email, u.name, u.role
		FROM events e
		LEFT JOIN users u ON e.owner_id = u.id
		WHERE e.id = $1
	`

	var event Event
	var owner User

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location,
		&owner.Id, &owner.Email, &owner.Name, &owner.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	event.Owner = &owner
	return &event, nil
}

func (m *EventModel) Update(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "UPDATE events SET name = $1, description = $2, date = $3, location = $4 WHERE id = $5"

	_, err := m.DB.ExecContext(ctx, query, event.Name, event.Description, event.Date, event.Location, event.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *EventModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "DELETE FROM events WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
