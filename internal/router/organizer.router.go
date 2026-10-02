package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/middleware"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func OrganizerRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	organizerRoute := router.Group("/organizer")

	or := repo.NewOrganizerRepo(db)
	os := service.NewOrganizerService(or)
	oh := handler.NewOrganizerHandler(os)

	organizerRoute.GET("",middleware.CheckToken(rc), middleware.OrganizerAccess, oh.GetDataDashboard)
}