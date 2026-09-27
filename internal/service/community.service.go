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

func (c *CommunityService) GetCommunity(ctx context.Context, communityId int) ([]dto.Community, error) {
	res, err := c.cr.GetCommunity(ctx, communityId)
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

func (c *CommunityService) GetCommunityBySearchFilter(ctx context.Context, search string, filter string) ([]dto.Community, error) {
	res, err := c.cr.GetCommunityBySearchFilter(ctx, search, filter)

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

func (c *CommunityService) GetCommunityMembers(ctx context.Context, communityId int) ([]dto.CommunityMembers, error) {
	res, err := c.cr.GetCommunityMembers(ctx, communityId)
	if err != nil {
		return nil, err
	}

	data := make([]dto.CommunityMembers, 0, len(res))
	for _, v := range res {
		data = append(data, dto.CommunityMembers{
			MemberName: v.Name,
		})
	}

	return data, nil
}
