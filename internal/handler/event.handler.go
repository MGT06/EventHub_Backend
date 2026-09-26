package handler

import (
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		Data:    res,
		Message: "Success Get Events",
	})

}

func (e *EventHandler) JoinEvent(ctx *gin.Context) {
	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	var body dto.JoinEvent
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	err := e.es.JoinEvent(ctx, idUser.(int), body.Id_Event)
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
		Message: "Success Join Event",
	})
}
