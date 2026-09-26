package repo

import (
	"context"
	"fmt"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (e *EventRepo) GetAllEvents(ctx context.Context) ([]model.EventDetail, error) {
	query := `SELECT e.id,
	   a.name,
       e.title,
       e.description,
       e.image_event_url,
       e.start_at,
       e.end_at,
       e.format,
       le.city,
       e.capacity,
       e.speakers,
       STRING_AGG(c.category_name, ', ') AS "category"
FROM events e
JOIN event_categories  ON event_categories.event_id = e.id
JOIN categories c ON event_categories.category_id = c.id
JOIN accounts a ON e.organizer_id = a.id
JOIN location_event le ON e.location_event_id = le.id
GROUP BY e.id, a.name, e.title, e.description, e.image_event_url, e.start_at, e.end_at, e.format,
         le.city, e.capacity, e.speakers;`

	res, err := e.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var events []model.EventDetail

	for res.Next() {
		var event model.EventDetail
		if err := res.Scan(&event.Id_event, &event.Name, &event.Title, &event.Description, &event.Image_event_url, &event.Start_at, &event.End_at, &event.Format, &event.City, &event.Capacity, &event.Speakers, &event.Category_name); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if res.Err() != nil {
		return nil, res.Err()
	}

	return events, nil
}

func (e *EventRepo) JoinEvent(ctx context.Context, idUser int, idEvent int) error {
	query := "INSERT INTO join_event (account_id, event_id) VALUES ($1, $2)"
	args := []any{idUser, idEvent}

	cmt, err := e.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}

	return nil
}

func (e *EventRepo) GetUpComingEvents(ctx context.Context) ([]model.EventDetail, error) {
	query := `SELECT e.id,
	   a.name,
       e.title,
       e.description,
       e.image_event_url,
       e.start_at,
       e.end_at,
       e.format,
       le.city,
       e.capacity,
       e.speakers,
       STRING_AGG(c.category_name, ', ') AS "category"
FROM events e
JOIN event_categories  ON event_categories.event_id = e.id
JOIN categories c ON event_categories.category_id = c.id
JOIN accounts a ON e.organizer_id = a.id
JOIN location_event le ON e.location_event_id = le.id
WHERE e.start_At > NOW()
GROUP BY e.id, a.name, e.title, e.description, e.image_event_url, e.start_at, e.end_at, e.format,
         le.city, e.capacity, e.speakers;`

	res, err := e.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var events []model.EventDetail

	for res.Next() {
		var event model.EventDetail
		if err := res.Scan(&event.Id_event, &event.Name, &event.Title, &event.Description, &event.Image_event_url, &event.Start_at, &event.End_at, &event.Format, &event.City, &event.Capacity, &event.Speakers, &event.Category_name); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if res.Err() != nil {
		return nil, res.Err()
	}

	return events, nil
}

func (e *EventRepo) GetMyEvent(ctx context.Context, idUser int) ([]model.EventDetail, error) {
	query := `SELECT e.id,
	   a.name,
       e.title,
       e.description,
       e.image_event_url,
       e.start_at,
       e.end_at,
       e.format,
       le.city,
       e.capacity,
       e.speakers,
       STRING_AGG(c.category_name, ', ') AS "category"
FROM events e
JOIN event_categories  ON event_categories.event_id = e.id
JOIN categories c ON event_categories.category_id = c.id
JOIN accounts a ON e.organizer_id = a.id
JOIN location_event le ON e.location_event_id = le.id
JOIN join_event je ON je.event_id = e.id
WHERE je.account_id = $1
GROUP BY e.id, a.name, e.title, e.description, e.image_event_url, e.start_at, e.end_at, e.format,
         le.city, e.capacity, e.speakers;`
	
	args := []any{idUser}

	res, err := e.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var events []model.EventDetail

	for res.Next() {
		var event model.EventDetail
		if err := res.Scan(&event.Id_event, &event.Name, &event.Title, &event.Description, &event.Image_event_url, &event.Start_at, &event.End_at, &event.Format, &event.City, &event.Capacity, &event.Speakers, &event.Category_name); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if res.Err() != nil {
		return nil, res.Err()
	}

	return events, nil
}