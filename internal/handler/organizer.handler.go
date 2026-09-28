package handler

import (
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type OrganizerHandler struct {
	os *service.OrganizerService
}

func NewOrganizerHandler(os *service.OrganizerService) *OrganizerHandler {
	return &OrganizerHandler{
		os: os,
	}
}

func (o *OrganizerHandler) GetDataDashboard(ctx *gin.Context) {
	idUser, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := o.os.GetDataDashboard(ctx.Request.Context(), idUser.(int))
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
		Message: "Success Get Data Dashboard",
	})
}
