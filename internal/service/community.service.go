package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type CommunityService struct {
	cr *repo.CommunityRepo
}

func NewCommunityService(cr *repo.CommunityRepo) *CommunityService {
	return &CommunityService{
		cr: cr,
	}
}

func (c *CommunityService) GetAllCommunity(ctx context.Context) ([]dto.Community, error) {
	res, err := c.cr.GetAllCommunity(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Community, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Community{
			Id:             v.Id_community,
			Community_name: v.Community_name,
			Description:    v.Description,
			Image_url:      v.Image_community_url,
			Category:       v.Category_name,
		})
	}

	return data, nil
}
