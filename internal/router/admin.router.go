package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/middleware"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AdminRouter(router *gin.Engine, db *pgxpool.Pool) {
	adminRoute := router.Group("/admin")

	ar := repo.NewAdminRepo(db)
	as := service.NewAdminService(ar)
	ah := handler.NewAdminHandler(as)

	adminRoute.GET("",middleware.CheckToken, middleware.AdminAccess, ah.GetDataDashboard)
}