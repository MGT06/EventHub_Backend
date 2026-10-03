package service

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventService struct {
	er *repo.EventRepo
	db *pgxpool.Pool
}

func NewEventService(er *repo.EventRepo, db *pgxpool.Pool) *EventService {
	return &EventService{
		er: er,
		db: db,
	}
}

func (e *EventService) GetEvents(ctx context.Context, eventId int) ([]dto.Event, error) {
	res, err := e.er.GetEvents(ctx, e.db, eventId)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		category := strings.Split(v.Category_name, ", ")

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
			Categories:    category,
		})
	}

	return data, nil
}

func (e *EventService) GetEventBySearchFilter(ctx context.Context, search string, filter string) ([]dto.Event, error) {
	res, err := e.er.GetEventBySearchFilter(ctx, e.db, search, filter)

	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		category := strings.Split(v.Category_name, ", ")

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
			Categories:    category,
		})
	}

	return data, nil
}

func (e *EventService) ToggleJoinEvent(ctx context.Context, idUser int, idEvent int) (bool, error) {
	isJoin, err := e.er.IsJoin(ctx, e.db, idUser, idEvent)
	if err != nil {
		return false, err
	}

	if isJoin {
		if err := e.er.LeaveEvent(ctx, e.db, idUser, idEvent); err != nil {
			return false, err
		}
		return isJoin, nil
	}

	if err := e.er.JoinEvent(ctx, e.db, idUser, idEvent); err != nil {
		return false, err
	}

	return isJoin, nil
}

func (e *EventService) GetUpComingEvents(ctx context.Context) ([]dto.Event, error) {
	res, err := e.er.GetUpComingEvents(ctx, e.db)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		category := strings.Split(v.Category_name, ", ")

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
			Categories:    category,
		})
	}

	return data, nil
}

func (e *EventService) GetMyEvent(ctx context.Context, idUser int) ([]dto.Event, error) {
	res, err := e.er.GetMyEvent(ctx, e.db, idUser)
	if err != nil {
		return nil, err
	}

	data := make([]dto.Event, 0, len(res))
	for _, v := range res {
		category := strings.Split(v.Category_name, ", ")

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
			Categories:    category,
		})
	}

	return data, nil
}

func (e *EventService) AddEvent(ctx context.Context, body dto.AddEvent, userId int, imgPath string,  speakers []dto.Speaker) error {
	tx, err := e.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Println(err.Error())
		}
	}()

	res, err := json.Marshal(speakers)
	if err != nil {
		return err
	}

	eventId, err := e.er.AddEvent(ctx, tx, model.Event{
		Organizer_id: userId,
		Community_id: body.CommunityId,
		Title: body.Title,
		Description: body.Description,
		Image_event_url: imgPath,
		Start_at: body.Start_at,
		End_at: body.End_at,
		Format: body.Format,
		Location_event_id: body.LocationId,
		Capacity: body.Capacity,
		Speakers: string(res),
	})
	if err != nil {
		return err
	}

	if err := e.er.AddCategory(ctx, tx, eventId, body.Categories); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
