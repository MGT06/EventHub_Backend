package repo

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TestimonyRepo struct {
	db *pgxpool.Pool
}

func NewTestimonyRepo(db *pgxpool.Pool) *TestimonyRepo {
	return &TestimonyRepo{
		db: db,
	}
}

func (t *TestimonyRepo) GetTestimony(ctx context.Context) ([]model.TestimonyDetail, error) {
	query := "SELECT a.name, t.message FROM testimony t JOIN accounts a ON t.account_id = a.id ORDER BY t.created_at ASC LIMIT 3"

	res, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var testimonies []model.TestimonyDetail
	for res.Next() {
		var testimony model.TestimonyDetail
		if err := res.Scan(&testimony.Name, &testimony.Message); err != nil {
			return nil, err
		}
		testimonies = append(testimonies, testimony)
	}

	return testimonies, nil
}