package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/middleware"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserRouter(router *gin.Engine, db *pgxpool.Pool) {
	userRoute := router.Group("/user")

	ur := repo.NewUserRepo(db)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)

	userRoute.GET("", middleware.CheckToken, uh.GetProfileUser)
}