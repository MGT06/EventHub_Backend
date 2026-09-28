package repo

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepo struct {
	db *pgxpool.Pool
}

func NewAdminRepo(db *pgxpool.Pool) *AdminRepo {
	return &AdminRepo{
		db: db,
	}
}

func (a *AdminRepo) GetDataDashboard(ctx context.Context) (model.DashboardAdmin, error) {
	query := `SELECT
    (SELECT COUNT(*)
     FROM accounts) AS "total_users",

    (SELECT COUNT(*)
     FROM events) AS "total_events",

    (SELECT COUNT(*)
     FROM communities) AS "total_communities"`

	res := a.db.QueryRow(ctx, query)

	var data model.DashboardAdmin
	if err := res.Scan(&data.TotalUsers, &data.TotalEvents, &data.TotalCommunities); err != nil {
		return model.DashboardAdmin{}, err
	}

	return data, nil
}
