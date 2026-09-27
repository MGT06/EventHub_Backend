package router

import (
	"github.com/MGT06/EventHub_Backend.git/internal/handler"
	"github.com/MGT06/EventHub_Backend.git/internal/middleware"
	"github.com/MGT06/EventHub_Backend.git/internal/repo"
	"github.com/MGT06/EventHub_Backend.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CommunityRouter(ctx *gin.Engine, db *pgxpool.Pool) {
	communityRoute := ctx.Group("/community")

	cr := repo.NewCommunityRepo(db)
	cs := service.NewCommunityService(cr)
	ch := handler.NewCommunityHandler(cs)

	communityRoute.GET("", middleware.CheckToken, ch.GetAllCommunity)
}
