package service

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
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

func (e *EventService) GetEvents(ctx context.Context, eventId int) ([]dto.Event, error) {
	res, err := e.er.GetEvents(ctx, eventId)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Event{
			Id:            v.Id_event,
			OrganizerName: v.Name,
			Title:         v.Title,
			Description:   v.Description,
			Image_url:     v.Image_event_url,
			Format:        v.Format,
			Location:      v.City,
			Capacity:      v.Capacity,
			Speakers:      v.Speakers,
			Categories:    v.Category_name,
		})
	}

	return data, nil
}

func (e *EventService) GetEventBySearchFilter(ctx context.Context, search string, filter string) ([]dto.Event, error) {
	res, err := e.er.GetEventBySearchFilter(ctx, search, filter)

	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Event{
			Id:            v.Id_event,
			OrganizerName: v.Name,
			Title:         v.Title,
			Description:   v.Description,
			Image_url:     v.Image_event_url,
			Format:        v.Format,
			Location:      v.City,
			Capacity:      v.Capacity,
			Speakers:      v.Speakers,
			Categories:    v.Category_name,
		})
	}

	return data, nil
}

func (e *EventService) ToggleJoinEvent(ctx context.Context, idUser int, idEvent int) (bool, error) {
	isJoin, err := e.er.IsJoin(ctx, idUser, idEvent)
	if err != nil {
		return false, err
	}

	if isJoin {
		if err := e.er.LeaveEvent(ctx, idUser, idEvent); err != nil {
			return false, err
		}
		return isJoin, nil
	}

	if err := e.er.JoinEvent(ctx, idUser, idEvent); err != nil {
		return false, err
	}

	return isJoin, nil
}

func (e *EventService) GetUpComingEvents(ctx context.Context) ([]dto.Event, error) {
	res, err := e.er.GetUpComingEvents(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Event{
			Id:            v.Id_event,
			OrganizerName: v.Name,
			Title:         v.Title,
			Description:   v.Description,
			Image_url:     v.Image_event_url,
			Format:        v.Format,
			Location:      v.City,
			Capacity:      v.Capacity,
			Speakers:      v.Speakers,
			Categories:    v.Category_name,
		})
	}

	return data, nil
}

func (e *EventService) GetMyEvent(ctx context.Context, idUser int) ([]dto.Event, error) {
	res, err := e.er.GetMyEvent(ctx, idUser)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		data = append(data, dto.Event{
			Id:            v.Id_event,
			OrganizerName: v.Name,
			Title:         v.Title,
			Description:   v.Description,
			Image_url:     v.Image_event_url,
			Format:        v.Format,
			Location:      v.City,
			Capacity:      v.Capacity,
			Speakers:      v.Speakers,
			Categories:    v.Category_name,
		})
	}

	return data, nil
}
