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

func EventRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	eventRoute := router.Group("/event")

	er := repo.NewEventRepo()
	es := service.NewEventService(er, db)
	eh := handler.NewEventHandler(es)

	eventRoute.GET("", eh.GetEventBySearchFilter)
	eventRoute.GET("detail/:id", eh.GetEvents)
	eventRoute.POST("join/:id", middleware.CheckToken(rc), eh.ToggleJoinEvent)
	eventRoute.POST("save/:id", middleware.CheckToken(rc), eh.ToggleSavedEvent)
	eventRoute.GET("upcoming", eh.GetUpComingEvents)
	eventRoute.GET("myevent", middleware.CheckToken(rc), eh.GetMyEvent)
	eventRoute.POST("create", middleware.CheckToken(rc), middleware.OrganizerAccess, eh.AddEvent)
}