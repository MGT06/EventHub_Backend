package handler

import (
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type CommunityHandler struct {
	cs *service.CommunityService
}

func NewCommunityHandler(cs *service.CommunityService) *CommunityHandler {
	return &CommunityHandler{
		cs: cs,
	}
}

func (c *CommunityHandler) GetAllCommunity(ctx *gin.Context) {
	res, err := c.cs.GetAllCommunity(ctx)
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
		Message: "Success Get Community",
	})

}
