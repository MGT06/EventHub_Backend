package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type NotificationService struct {
	nr *repo.NotificationRepo
}

func NewNotificationService(nr *repo.NotificationRepo) *NotificationService {
	return &NotificationService{
		nr: nr,
	}
}

func (n *NotificationService) GetMyNotification(ctx context.Context, userId int) ([]dto.Notification, error) {
	res, err := n.nr.GetMyNotification(ctx, userId)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Notification, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Notification{
			Id: v.NotifId,
			UserName: v.Name,
			Title: v.Title,
			Message: v.Message,
			Type: v.Type,
			Read_at: v.Read_at,
		})
	}

	return data, nil
}
