package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type UserService struct {
	ur *repo.UserRepo
}

func NewUserService(ur *repo.UserRepo) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (u *UserService) GetProfileUser(ctx context.Context, userId int) (dto.UserProfile, error) {
	res, err := u.ur.GetProfileUser(ctx, userId)

	if err != nil {
		return dto.UserProfile{}, err
	}

	data := dto.UserProfile{
		Name:          res.Name,
		Bio:           res.Bio,
		User_location: res.User_location,
		Position:      res.Position,
		Avatar_url:    res.Avatar_url,
	}

	return data, err
}

func (u *UserService) EditProfileUser(ctx context.Context, body dto.UserProfile, userId int) error {
	if err := u.ur.EditProfileUser(ctx, model.Account{
		Name:          body.Name,
		Bio:           body.Bio,
		User_location: body.User_location,
		Position:      body.Position,
		Avatar_url:    body.Avatar_url,
	}, userId); err != nil {
		return err
	}

	return nil
}
