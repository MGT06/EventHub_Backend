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
	eventRoute := router.Group("/event")

	er := repo.NewEventRepo(db)
	es := service.NewEventService(er)
	eh := handler.NewEventHandler(es)

	eventRoute.GET("", middleware.CheckToken, eh.GetEventBySearchFilter)
	eventRoute.GET("detail/:id", middleware.CheckToken, eh.GetEvents)
	eventRoute.POST("join", middleware.CheckToken, eh.JoinEvent)
	eventRoute.GET("upcoming", middleware.CheckToken, eh.GetUpComingEvents)
	eventRoute.GET("myevent", middleware.CheckToken, eh.GetMyEvent)
}
