package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

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

// Get Event
//
// @Summary			Get Event
// @Description		Get event detail using ID event
// @Tags			event
// @Produce			json
// @Router			/event/detail/{id}	[get]
// @Param			id		path		int		true	"Event ID"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.Response
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

// Toggle Join Event
//
// @Summary			Join and Leave Event
// @Description		Toggle to Join and Leave Event
// @Tags			event
// @Produce			json
// @Router			/event/join/{id}		[post]
// @Security 		BearerToken
// @Param			id		path	int	true	"Event ID"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.Response
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

func (e *EventHandler) ToggleSavedEvent(ctx *gin.Context) {
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

	isSaved, er := e.es.ToggleSavedEvent(ctx.Request.Context(), idUser.(int), idEvent)
	if er != nil {
		log.Println(er)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	if isSaved {
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Message: "Success UnSaved Event",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Saved Event",
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

// Create Event
//
// @Summary			Create New Event
// @Tags			event
// @Accept			mpfd
// @Produce			json
// @Router			/event/create	[post]
// @Security 		BearerToken
// @Param			community_id			formData	int			false	"add community to event"
// @Param			location_event_id		formData	int			false	"add location event"
// @Param			title					formData	string		false	"add title event"
// @Param			description				formData	string		false	"add desc event"
// @Param			image					formData	file		false	"add image event"
// @Param			start_at				formData	string		false	"add time start event" format(date-time)
// @Param			end_at					formData	string		false	"add time end event" format(date-time)
// @Param			format					formData	string		false	"add format event"
// @Param			capacity				formData	string		false	"add capacity event"
// @Param			speakers				formData	string		false	"add speakers event"
// @Param			categories				formData	[]int		false	"add categories event" collectionFormat(multi)
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.Response
func (e *EventHandler) AddEvent(ctx *gin.Context) {
	idUser, exist := ctx.Get("idUser")
	if !exist {
		log.Println(idUser)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	var body dto.AddEvent
	if err := ctx.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), body.Title, path.Ext(body.Image.Filename))
	filepath := path.Join("public", "img", "events", filename)

	if err := ctx.SaveUploadedFile(&body.Image, filepath); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	var speakers []dto.Speaker
	if err := json.Unmarshal([]byte(body.Speakers), &speakers); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid speakers format",
		})
		return
	}

	if err := e.es.AddEvent(ctx.Request.Context(), body, idUser.(int), filepath, speakers); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Create Event",
	})
}
