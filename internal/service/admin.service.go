package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type AdminService struct {
	ar *repo.AdminRepo
}

func NewAdminService(ar *repo.AdminRepo) *AdminService {
	return &AdminService{
		ar: ar,
	}
}

func (a *AdminService) GetDataDashboard(ctx context.Context) (dto.DashboardAdmin, error) {
	res, err := a.ar.GetDataDashboard(ctx)
	if err != nil {
		return dto.DashboardAdmin{}, err
	}

	data := dto.DashboardAdmin{
		TotalUsers: res.TotalUsers,
		TotalEvents: res.TotalEvents,
		TotalCommunities: res.TotalCommunities,
	}

	return data, nil
}