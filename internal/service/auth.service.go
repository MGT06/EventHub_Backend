package service

import (
	"context"
	"errors"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	errorTemplate "github.com/MGT06/EventHub_Backend.git/internal/error"
	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/pkg"
	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	ar *repo.AuthRepo
}

func NewAuthService(ar *repo.AuthRepo) *AuthService {
	return &AuthService{
		ar: ar,
	}
}

func (a *AuthService) Register(ctx context.Context, body dto.Register) error {
	if len(body.FullName) == 0 || len(body.Email) == 0 || len(body.Password) == 0 {
		return errorTemplate.ErrInvalidInputs
	}

	_, err := a.ar.FindAccount(ctx, body.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	hash := pkg.NewHashConfig().GenHash(body.Password)

	if err := a.ar.Register(ctx, model.Account{
		Name:     body.FullName,
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Login(ctx context.Context, body dto.Login) (string, error) {
	if len(body.Email) == 0 || len(body.Password) == 0 {
		return "",  errorTemplate.ErrInvalidInputs
	}

	acc, err := a.ar.FindAccount(ctx, body.Email)
	if err != nil {
		return "", errorTemplate.ErrEmailPasswordIncorrect
	}

	if err := pkg.Compare(body.Password, acc.Password); err != nil {
		return "", errorTemplate.ErrEmailPasswordIncorrect
	}

	claims := pkg.NewJWTClaims(acc.Id, acc.Role)
	return claims.GenToken()
}
