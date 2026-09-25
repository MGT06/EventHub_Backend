package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func authRouter(router *gin.Engine, db *pgxpool.Pool){
	authRoute := router.Group("/auth")

	ar := repo.NewAuthRepo(db)
	as := service.NewAuthService(ar)
	ah := handler.NewAuthHandler(as)

	authRoute.POST("/register", ah.Register)
	authRoute.POST("/login", ah.Login)
}