package handler

import (
	"log"
	"net/http"
	"strconv"

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

func (e *EventHandler) GetEvents(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := e.es.GetEvents(ctx.Request.Context(), eventId)

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

func (e *EventHandler) GetEventBySearchFilter(ctx *gin.Context) {

	search := ctx.Query("search")
	filter := ctx.Query("category")

	res, err := e.es.GetEventBySearchFilter(ctx.Request.Context(), search, filter)

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

func (e *EventHandler) ToggleJoinEvent(ctx *gin.Context) {
	idEvent, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}
	isJoin, er := e.es.ToggleJoinEvent(ctx.Request.Context(), idUser.(int), idEvent)
	if er != nil {
		log.Println(er)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	if isJoin {
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Message: "Success Leave Event",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Join Event",
	})
}

func (e *EventHandler) GetUpComingEvents(ctx *gin.Context) {
	res, err := e.es.GetUpComingEvents(ctx.Request.Context())
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

func (e *EventHandler) GetMyEvent(ctx *gin.Context) {
	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := e.es.GetMyEvent(ctx.Request.Context(), idUser.(int))
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
