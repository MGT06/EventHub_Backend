package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	errorTemplate "github.com/MGT06/EventHub_Backend.git/internal/error"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TestimonyHandler struct {
	ts *service.TestimonyService
}

func NewTestimonyHandler(ts *service.TestimonyService) *TestimonyHandler {
	return &TestimonyHandler{
		ts: ts,
	}
}

// Get Testimony
//
// @Summary			Get Testimony App
// @Tags			testimony
// @Produce			json
// @Router			/testimony	[get]
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.ErrorResponse
func (t *TestimonyHandler) GetTestimony(ctx *gin.Context) {
	res, err := t.ts.GetTestimony(ctx.Request.Context())

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
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

// Set Testimony
//
// @Summary			Set Testimony App
// @Tags			testimony
// @Accept			json
// @Produce			json
// @Router			/testimony/send-testimony	[post]
// @Param			newTestimony	body	dto.SetTestimony	true	"Set Testimony App"
// @Security 		BearerToken
// @Success			200		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			401		{object}	dto.ErrorResponse
// @Failure			409		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (t *TestimonyHandler) SetTestimony(ctx *gin.Context) {
	userId, exist := ctx.Get("idUser")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	var newTestimony dto.SetTestimony
	if err := ctx.ShouldBindWith(&newTestimony, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	err := t.ts.SetTestimony(ctx.Request.Context(),newTestimony.Message, userId.(int))
	if err != nil {
		log.Println(err)
		if errors.Is(err, errorTemplate.ErrInvalidInputs) {
			ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Success: false,
				Message: "Please fill in all required fields",
			})
			return
		}
		if errors.Is(err, errorTemplate.ErrDoubleSubmittedTestimony) {
			ctx.JSON(http.StatusConflict, dto.ErrorResponse{
				Success: false,
				Message: errorTemplate.ErrDoubleSubmittedTestimony.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: true,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Set Testimony",
	})
}
