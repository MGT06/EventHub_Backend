package repo

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizerRepo struct {
	db *pgxpool.Pool
}

func NewOrganizerRepo(db *pgxpool.Pool) *OrganizerRepo {
	return &OrganizerRepo{
		db: db,
	}
}

func (o *OrganizerRepo) GetDataDashboard(ctx context.Context, userId int) (model.DashboardOrganizer, error) {
	query := `SELECT
    (SELECT COUNT(title)
     FROM events
     WHERE organizer_id = $1) AS "total_event_created",

    (SELECT COUNT(DISTINCT je.account_id)
     FROM join_event je
     JOIN events e ON je.event_id = e.id
     WHERE e.organizer_id = $1) AS "total_attendees_joined",

    (SELECT COALESCE(COUNT(DISTINCT je.account_id) * 100 / NULLIF(COUNT(e.id), 0), 0)
     FROM join_event je
     JOIN events e ON je.event_id = e.id
     WHERE e.organizer_id = $1) AS "avg_fill_rate";`

	args := []any{userId}

	res := o.db.QueryRow(ctx, query, args...)

	var data model.DashboardOrganizer
	if err := res.Scan(&data.TotalEvent, &data.TotalAttendes, &data.AVGFillRate); err != nil {
		return model.DashboardOrganizer{}, err
	}

	return data, nil
}
