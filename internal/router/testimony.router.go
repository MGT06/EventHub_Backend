package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestimonyRouter(router *gin.Engine, db *pgxpool.Pool) {
	testimonyRoute := router.Group("/testimony")

	tr := repo.NewTestimonyRepo(db)
	ts := service.NewTestimonyService(tr)
	th := handler.NewTestimonyHandler(ts)

	testimonyRoute.GET("", th.GetTestimony)
}