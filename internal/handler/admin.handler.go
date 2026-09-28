package handler

import (
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	as *service.AdminService
}

func NewAdminHandler(as *service.AdminService) *AdminHandler {
	return &AdminHandler{
		as: as,
	}
}

func (a *AdminHandler) GetDataDashboard(ctx *gin.Context) {
	res, err := a.as.GetDataDashboard(ctx.Request.Context())
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
