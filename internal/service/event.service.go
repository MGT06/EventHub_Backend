package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	errorTemplate "github.com/MGT06/EventHub_Backend.git/internal/error"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
)

type EventService struct {
	er *repo.EventRepo
}

func NewEventService(er *repo.EventRepo) *EventService {
	return &EventService{
		er: er,
	}
}

func (e *EventService) GetAllEvents(ctx context.Context) ([]dto.Event, error) {
	res, err := e.er.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Event{
			Id: v.Id_event,
			OrganizerName: v.Name,
			Title: v.Title,
			Description: v.Description,
			Image_url: v.Image_event_url,
			Format: v.Format,
			Location: v.City,
			Capacity: v.Capacity,
			Speakers: v.Speakers,
			Categories: v.Category_name,
		})
	}

	return data, nil
}

func (e *EventService) JoinEvent(ctx context.Context, idUser int, idEvent int) error {
	if idUser == 0 || idEvent == 0 {
		return errorTemplate.ErrInvalidInputs
	}

	if err := e.er.JoinEvent(ctx, idUser, idEvent); err != nil {
		return err
	}

	return nil
}