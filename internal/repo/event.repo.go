package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type EventRepo struct{}

func NewEventRepo() *EventRepo {
	return &EventRepo{}
}

func (e *EventRepo) GetEvents(ctx context.Context, db DBTX, eventId int) ([]model.EventDetail, error) {
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
WHERE e.id = $1
GROUP BY e.id, a.name, e.title, e.description, e.image_event_url, e.start_at, e.end_at, e.format,
         le.city, e.capacity, e.speakers;`

	args := []any{eventId}

	res, err := db.Query(ctx, query, args...)
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

func (e *EventRepo) GetEventBySearchFilter(ctx context.Context, db DBTX, search string, filter string) ([]model.EventDetail, error) {
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
WHERE e.title ILIKE $1
GROUP BY e.id, a.name, e.title, e.description, e.image_event_url, e.start_at, e.end_at, e.format,
         le.city, e.capacity, e.speakers
HAVING STRING_AGG(c.category_name, ', ') ILIKE $2;`

	args := []any{"%" + search + "%", "%" + filter + "%"}

	res, err := db.Query(ctx, query, args...)
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

func (e *EventRepo) JoinEvent(ctx context.Context, db DBTX, idUser int, idEvent int) error {
	query := "INSERT INTO join_event (account_id, event_id) VALUES ($1, $2)"
	args := []any{idUser, idEvent}

	cmt, err := db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}

	return nil
}

func (e *EventRepo) LeaveEvent(ctx context.Context, db DBTX, idUser int, idEvent int) error {
	query := "DELETE FROM join_event WHERE account_id = $1 AND event_id = $2"
	args := []any{idUser, idEvent}

	cmt, err := db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}

	return nil
}

func (e *EventRepo) IsJoin(ctx context.Context, db DBTX, idUser int, idEvent int) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM join_event WHERE account_id = $1 AND event_id = $2)"
	args := []any{idUser, idEvent}

	res := db.QueryRow(ctx, query, args...)

	var isJoin bool
	if err := res.Scan(&isJoin); err != nil {
		return false, err
	}

	return isJoin, nil
}

func (e *EventRepo) GetUpComingEvents(ctx context.Context, db DBTX) ([]model.EventDetail, error) {
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

	res, err := db.Query(ctx, query)
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

func (e *EventRepo) GetMyEvent(ctx context.Context, db DBTX, idUser int) ([]model.EventDetail, error) {
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

	res, err := db.Query(ctx, query, args...)
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

func (e *EventRepo) AddEvent(ctx context.Context, db DBTX, body model.Event) (int, error) {
	query := "INSERT INTO events (organizer_id, community_id, location_event_id, title, description, image_event_url, start_at, end_at, format, capacity, speakers) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"

	args := []any{body.Organizer_id, body.Community_id, body.Location_event_id, body.Title, body.Description, body.Image_event_url, body.Start_at, body.End_at, body.Format, body.Capacity, body.Speakers}

	var eventId int
	if err := db.QueryRow(ctx, query, args...).Scan(&eventId); err != nil {
		return 0, err
	}

	return eventId, nil
}

func (e *EventRepo) AddCategory(ctx context.Context, db DBTX, eventId int, categoryId []int) error {
	if len(categoryId) == 0 {
		return nil
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO event_categories (event_id, category_id) VALUES")

	args := make([]any, 0, len(categoryId)*2)
	SQLParam := make([]string, 0, len(categoryId))

	for idx, category := range categoryId {
		SQLParam = append(SQLParam, fmt.Sprintf("($%d, $%d)", (idx*2)+1, (idx*2)+2))
		args = append(args, eventId, category)
	}

	sb.WriteString(strings.Join(SQLParam, ", "))

	cmt, err := db.Exec(ctx, sb.String(), args...)
	if err != nil {
		return err
	}
	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}


	return nil
}
