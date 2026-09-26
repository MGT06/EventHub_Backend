package handler

import (
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	es *service.EventService
}

func NewEventHandler(es *service.EventService) *EventHandler {
	return &EventHandler{
		es: es,
	}
}

func (e *EventHandler) GetAllEvents(ctx *gin.Context) {
	res, err := e.es.GetAllEvents(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: res,
		Message: "Success Get Events",
	})
	
}