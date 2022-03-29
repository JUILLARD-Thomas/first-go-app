package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "user"
)

type Event struct {
	Type_enum  string  `json:"type"`
	User_agent string  `json:"user_agent"`
	Ip         string  `json:"ip"`
	Ts         float64 `json:"timestamp"`
}

func Events(db *sql.DB) ([]Event, error) {
	rows, err := db.Query("SELECT * FROM event")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An Event slice to hold data from returned rows.
	var events []Event

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.User_agent, &event.Type_enum, &event.Ip,
			&event.Ts, &event); err != nil {
			return events, err
		}
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return events, err
	}
	return events, nil
}
