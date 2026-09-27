package handler

import (
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type TestimonyHandler struct {
	ts *service.TestimonyService
}

func NewTestimonyHandler(ts *service.TestimonyService) *TestimonyHandler {
	return &TestimonyHandler{
		ts: ts,
	}
}

func (t *TestimonyHandler) GetTestimony(ctx *gin.Context) {
	res, err := t.ts.GetTestimony(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: true,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    res,
		Message: "Success Get Testimony",
	})
}
