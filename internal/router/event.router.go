package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/middleware"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func EventRouter(router *gin.Engine, db *pgxpool.Pool) {
	EventRoute := router.Group("/event")

	er := repo.NewEventRepo(db)
	es := service.NewEventService(er)
	eh := handler.NewEventHandler(es)

	EventRoute.GET("", middleware.CheckToken, eh.GetAllEvents)
}
