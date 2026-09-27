package repo

import (
	"context"
	"fmt"

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

func (t *TestimonyRepo) SetTestimony(ctx context.Context, message string, userId int) error {
	query := "INSERT INTO testimony (account_id, message) VALUES ($1, $2)"
	args := []any{userId, message}

	cmt, err := t.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}

	return nil
}

func (t *TestimonyRepo) CheckUserTestimony(ctx context.Context, userId int) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM testimony WHERE account_id = $1)"
	args := []any{userId}

	res := t.db.QueryRow(ctx, query, args...)

	var checkUserTestimony bool
	if err := res.Scan(&checkUserTestimony); err != nil {
		return false, err
	}

	return checkUserTestimony, nil
}
