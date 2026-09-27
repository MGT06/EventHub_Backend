package handler

import (
	"log"
	"net/http"
	"strconv"

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

func (c *CommunityHandler) GetCommunity(ctx *gin.Context) {
	communityId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := c.cs.GetCommunity(ctx.Request.Context(), communityId)
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

func (c *CommunityHandler) GetCommunityBySearchFilter(ctx *gin.Context) {
	search := ctx.Query("search")
	filter := ctx.Query("category")

	res, err := c.cs.GetCommunityBySearchFilter(ctx.Request.Context(), search, filter)
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

func (c *CommunityHandler) GetCommunityMembers(ctx *gin.Context) {
	communityId := ctx.Param("id")
	id, err := strconv.Atoi(communityId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := c.cs.GetCommunityMembers(ctx.Request.Context(), id)
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
