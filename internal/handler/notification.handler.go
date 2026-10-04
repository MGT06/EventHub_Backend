package handler

import (
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	ns *service.NotificationService
}

func NewNotificationHandler(ns *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		ns: ns,
	}
}

// Get User Notification
//
// @Summary			Get User Notification
// @Tags			notification
// @Produce			json
// @Router			/notification	[get]
// @Security 		BearerToken
// @Success			200		{object}	dto.Response
// @Failure			401		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (n *NotificationHandler) GetMyNotification(ctx *gin.Context) {
	id, exists := ctx.Get("idUser")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := n.ns.GetMyNotification(ctx.Request.Context(), id.(int))
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
		Message: "Success Get Notification",
	})
}
