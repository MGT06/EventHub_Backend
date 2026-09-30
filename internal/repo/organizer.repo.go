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
	query := `SELECT COUNT(e.id) AS "total_event_created",
       COUNT(DISTINCT je.account_id) AS "total_attendees_joined",
       COALESCE(COUNT(DISTINCT je.account_id) * 100 / NULLIF(COUNT(e.id), 0), 0) AS "avg_fill_rate" 
	   FROM events e
	   LEFT JOIN join_event je ON e.id = je.event_id
	   WHERE e.organizer_id = 3;`


	args := []any{userId}

	res := o.db.QueryRow(ctx, query, args...)

	var data model.DashboardOrganizer
	if err := res.Scan(&data.TotalEvent, &data.TotalAttendes, &data.AVGFillRate); err != nil {
		return model.DashboardOrganizer{}, err
	}

	return data, nil
}
