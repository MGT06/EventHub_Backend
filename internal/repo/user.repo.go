package repo

import (
	"context"
	"fmt"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) GetProfileUser(ctx context.Context, userId int) (model.Account, error) {
	query := "SELECT name, bio, user_location, position, avatar_url FROM accounts where id = $1"
	args := []any{userId}

	res := u.db.QueryRow(ctx,query, args...)

	var profile model.Account
	if err := res.Scan(&profile.Name, &profile.Bio, &profile.User_location, &profile.Position, &profile.Avatar_url); err != nil {
		return model.Account{}, nil
	}

	fmt.Println(profile.Name)

	return profile, nil
}
