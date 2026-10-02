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

func TestimonyRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	testimonyRoute := router.Group("/testimony")

	tr := repo.NewTestimonyRepo(db)
	ts := service.NewTestimonyService(tr)
	th := handler.NewTestimonyHandler(ts)

	testimonyRoute.GET("", th.GetTestimony)
	testimonyRoute.POST("/send-testimony", middleware.CheckToken(rc), th.SetTestimony)
}