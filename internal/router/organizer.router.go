package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/middleware"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OrganizerRouter(router *gin.Engine, db *pgxpool.Pool) {
	organizerRoute := router.Group("/organizer")

	or := repo.NewOrganizerRepo(db)
	os := service.NewOrganizerService(or)
	oh := handler.NewOrganizerHandler(os)

	organizerRoute.GET("",middleware.CheckToken, middleware.OrganizerAccess, oh.GetDataDashboard)
}