package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	authRouter(router, db)
	EventRouter(router, db)
	CommunityRouter(router, db)
	UserRouter(router, db, rc)
	TestimonyRouter(router, db)
	notificationRouter(router, db)
	OrganizerRouter(router, db)
	AdminRouter(router, db)

	router.GET("api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
