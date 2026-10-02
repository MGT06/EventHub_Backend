package handler

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (u *UserHandler) GetProfileUser(ctx *gin.Context) {
	id, exists := ctx.Get("idUser")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	res, err := u.us.GetProfileUser(ctx.Request.Context(), id.(int))

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
		Message: "Success Get Profile",
	})

}

// Edit User
//
// @Summary			Update user info
// @Tags			user
// @Accept			mpfd
// @Produce			json
// @Router			/user/edit	[patch]
// @Security 		BearerToken
// @Param			name			formData	string	true	"update name user"
// @Param			bio				formData	string	false	"update bio user"
// @Param			user_location	formData	string	false	"update location user"
// @Param			position		formData	string	false	"update position user"
// @Param			avatar_url		formData	file	false	"update avatar user"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.Response
func (u *UserHandler) EditProfileUser(ctx *gin.Context) {
	id, exists := ctx.Get("idUser")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	var body dto.SetUserProfile
	if err := ctx.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), body.Name, path.Ext(body.Avatar.Filename))
	filepath := path.Join("public", "img", filename)

	if err := ctx.SaveUploadedFile(body.Avatar, filepath); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	if err := u.us.EditProfileUser(ctx, body, id.(int), filepath); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Edit Profile",
	})

}
