package models

import (
	"errors"
	"goapi/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// add to a database
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, eventId)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetCreatedEvent(userId int64) ([]Event, error) {
	query := "SELECT * FROM events WHERE user_id = ?"
	row, err := db.DB.Query(query, userId)

	if err != nil {
		return nil, err
	}

	var myEvents []Event

	for row.Next() {
		var event Event
		err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		myEvents = append(myEvents, event)
	}	

	return myEvents, nil
}

func GetRegisteredEvent(userId int64) ([]Event, error) {
	query := "SELECT event_id FROM registrations WHERE user_id = ?"

	row, err := db.DB.Query(query, userId)

	if err != nil {
		return nil, err
	}


	var myEvents []Event

	for row.Next() {

		var eventId int64
		err = row.Scan(&eventId)
		if err != nil {
			return nil, err
		}

		event, err := GetEventById(eventId)
		if err != nil {
			return nil, err
		}

		myEvents = append(myEvents, *event)
	}	

	return myEvents, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func Deregister(userId, eventId int64) error {
	query := "DELETE FROM registrations WHERE user_id = ? AND event_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(userId, eventId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no recods found")
	}

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	if err != nil {
		return err
	}

	return nil
}
