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

type AuthHandler struct {
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
	}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var newAccount dto.Register
	if err := ctx.ShouldBindWith(&newAccount, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	if err := a.as.Register(ctx.Request.Context(), newAccount); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Account created successfully",
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var account dto.Login
	if err := ctx.ShouldBindWith(&account, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	token, err := a.as.Login(ctx.Request.Context(), account)
	if err != nil {
		log.Println(err)
		if errors.Is(err, errorTemplate.ErrInvalidInputs) {
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "Please fill in all required fields",
			})
			return
		}
		if errors.Is(err, errorTemplate.ErrEmailPasswordIncorrect) {
			ctx.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"token": token,
		},
		Message: "Login Success",
	})
}
