package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"id"`
	EventId int `json:"event_id"`
	UserId  int `json:"user_id"`
}
