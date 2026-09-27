package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
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
