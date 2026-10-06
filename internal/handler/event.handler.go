package handler

import (
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
// @Accept 			json
// @Produce			json
// @Router			/event/detail/{id}	[get]
// @Param 			id		path	int		true	"Event Id"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.ErrorResponse
func (e *EventHandler) GetEvents(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := e.es.GetEvents(ctx.Request.Context(), eventId)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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

// Get Event
//
// @Summary			Get Event By Search And Filter
// @Tags			event
// @Produce			json
// @Router			/event	[get]
// @Param			search			query		string		false	"search query param"
// @Param			category		query		string		false	"category query param"
// @Param			location		query		string		false	"location query param"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.ErrorResponse
func (e *EventHandler) GetEventBySearchFilter(ctx *gin.Context) {
	search := ctx.Query("search")
	category := ctx.Query("category")
	location := ctx.Query("location")

	res, err := e.es.GetEventBySearchFilter(ctx.Request.Context(), search, category, location)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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
// @Failure			500		{object}	dto.ErrorResponse
func (e *EventHandler) ToggleJoinEvent(ctx *gin.Context) {
	idEvent, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}
	isJoin, er := e.es.ToggleJoinEvent(ctx.Request.Context(), idUser.(int), idEvent)
	if er != nil {
		log.Println(er)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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

// Toggle Saved Event
//
// @Summary			Saved and Unsaved Event
// @Description		Toggle to Saved and Unsaved Event
// @Tags			event
// @Produce			json
// @Router			/event/save/{id}		[post]
// @Security 		BearerToken
// @Param			id		path	int	true	"Event ID"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.ErrorResponse
func (e *EventHandler) ToggleSavedEvent(ctx *gin.Context) {
	idEvent, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	isSaved, er := e.es.ToggleSavedEvent(ctx.Request.Context(), idUser.(int), idEvent)
	if er != nil {
		log.Println(er)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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

// Get Upcoming Events
//
// @Summary			Get Upcoming Events
// @Tags			event
// @Produce			json
// @Router			/event/upcoming	[get]
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.ErrorResponse
func (e *EventHandler) GetUpComingEvents(ctx *gin.Context) {
	res, err := e.es.GetUpComingEvents(ctx.Request.Context())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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

// Get Event
//
// @Summary			Get Event
// @Description		Get event detail using ID event
// @Tags			event
// @Produce			json
// @Router			/event/myevent	[get]
// @Security 		BearerToken
// @Success			200		{object}	dto.Response
// @Failure			401		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (e *EventHandler) GetMyEvent(ctx *gin.Context) {
	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := e.es.GetMyEvent(ctx.Request.Context(), idUser.(int))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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
// @Param			location_event_id		formData	int			true	"add location event"
// @Param			title					formData	string		true	"add title event"
// @Param			description				formData	string		true	"add desc event"
// @Param			image					formData	file		true	"add image event"
// @Param			start_at				formData	string		true	"add time start event" format(date-time)
// @Param			end_at					formData	string		true	"add time end event" format(date-time)
// @Param			format					formData	string		true	"add format event"
// @Param			capacity				formData	string		true	"add capacity event"
// @Param			speakers				formData	string		false	"add speakers event"
// @Param			categories				formData	[]int		true	"add categories event" collectionFormat(multi)
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

	ext := path.Ext(body.Image.Filename)
	switch ext {
		case ".jpg", ".jpeg", ".png", ".webp":
		default:
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "Unsupported image type",
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

	if err := e.es.AddEvent(ctx.Request.Context(), body, idUser.(int), filepath); err != nil {
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

// Edit Event
//
// @Summary			Edit Event
// @Tags			event
// @Accept			mpfd
// @Produce			json
// @Router			/event/{id}/edit	[patch]
// @Security 		BearerToken
// @Param			id						path		int			true	"event id"
// @Param			community_id			formData	int			false	"community id"
// @Param			location_event_id		formData	int			false	"location event id"
// @Param			title					formData	string		false	"title event"
// @Param			description				formData	string		false	"desc event"
// @Param			image					formData	file		false	"image event"
// @Param			start_at				formData	string		false	"time start event" format(date-time)
// @Param			end_at					formData	string		false	"time end event" format(date-time)
// @Param			format					formData	string		false	"format event"
// @Param			capacity				formData	int			false	"capacity event"
// @Param			speakers				formData	string		false	"speakers event (JSON)"
// @Param			categories				formData	[]int		false	"categories event" collectionFormat(multi)
// @Success			200		{object}	dto.Response
// @Failure			400		{object}	dto.Response
// @Failure			500		{object}	dto.Response
func (e *EventHandler) EditEvent(ctx *gin.Context) {
	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid event id",
		})
		return
	}

	var body dto.EditEvent
	if err := ctx.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}
	
	var filepath string
	if body.Image != nil {
		ext := path.Ext(body.Image.Filename)
		switch ext {
		case ".jpg", ".jpeg", ".png", ".webp":
		default:
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "Unsupported image type",
			})
			return
		}
	
		filename := fmt.Sprintf("%d_%d%s", time.Now().UnixNano(), eventId, ext)
		filepath = path.Join("public", "img", "events", filename)
	
		if err := ctx.SaveUploadedFile(body.Image, filepath); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "A system error has occurred",
			})
			return
		}
	}
	

	if err := e.es.EditEvent(ctx.Request.Context(), eventId, body, idUser.(int), filepath); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Update Event",
	})
}
