package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type TestimonyService struct {
	tr *repo.TestimonyRepo
}

func NewTestimonyService(tr *repo.TestimonyRepo) *TestimonyService {
	return &TestimonyService{
		tr: tr,
	}
}

func (t *TestimonyService) GetTestimony(ctx context.Context) ([]dto.Testimony, error) {
	res, err := t.tr.GetTestimony(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Testimony, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Testimony{
			UserName: v.Name,
			Message: v.Message,
		})
	}

	return data, nil
}
