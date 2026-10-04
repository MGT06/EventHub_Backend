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
	query := "SELECT name, email, bio, user_location, position, avatar_url FROM accounts where id = $1"
	args := []any{userId}

	res := u.db.QueryRow(ctx, query, args...)

	var profile model.Account
	if err := res.Scan(&profile.Name, &profile.Email, &profile.Bio, &profile.User_location, &profile.Position, &profile.Avatar_url); err != nil {
		return model.Account{}, nil
	}

	fmt.Println(profile.Name)

	return profile, nil
}

func (u *UserRepo) EditProfileUser(ctx context.Context, body model.Account, userId int) error {
	query := `UPDATE accounts 
    SET 
		name = COALESCE($1, name),
        bio = COALESCE($2, bio), 
        user_location = COALESCE($3, user_location), 
        position = COALESCE($4, position), 
        avatar_url = COALESCE($5, avatar_url) 
    WHERE id = $6`
	args := []any{body.Name, body.Bio, body.User_location, body.Position, body.Avatar_url, userId}

	cmt, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}

	return nil
}

func (u *UserRepo) GetCurrPassword(ctx context.Context, userId int) (string, error) {
	query := "SELECT password FROM accounts WHERE id = $1"
	args := []any{userId}

	res := u.db.QueryRow(ctx, query, args...)

	var currPassword string
	if err := res.Scan(&currPassword); err != nil {
		return "", err
	}

	return currPassword, nil
}

func (u *UserRepo) ChangePassword(ctx context.Context, userId int, newPassword string) error {
	query := "UPDATE accounts SET password = $1 WHERE id = $2"
	args := []any{newPassword, userId}

	cmt, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return fmt.Errorf("no row affected")
	}

	return nil
}

func (u *UserRepo) GetUserHeaderInformation(ctx context.Context, userId int) (model.Account, error) {
	query := "SELECT name, email, avatar_url FROM accounts WHERE id = $1"
	args := []any{userId}

	res := u.db.QueryRow(ctx, query, args...)
	
	var userInformation model.Account
	if err := res.Scan(&userInformation.Name, &userInformation.Email, &userInformation.Avatar_url); err != nil {
		return model.Account{}, err
	}

	return userInformation, nil
}
