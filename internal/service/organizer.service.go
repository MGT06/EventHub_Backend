package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type OrganizerService struct {
	or *repo.OrganizerRepo
}

func NewOrganizerService(or *repo.OrganizerRepo) *OrganizerService {
	return &OrganizerService{
		or: or,
	}
}


func (o *OrganizerService) GetDataDashboard(ctx context.Context, userId int) (dto.DashboardOrganizer, error) {
	res, err := o.or.GetDataDashboard(ctx, userId)
	if err != nil {
		return dto.DashboardOrganizer{}, err
	}

	data := dto.DashboardOrganizer{
		TotalEvent: res.TotalEvent,
		TotalAttendes: res.TotalAttendes,
		AVGFillRate: res.AVGFillRate,
	}

	return data, nil
}