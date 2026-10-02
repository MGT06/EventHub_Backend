package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	authRouter(router, db, rc)
	EventRouter(router, db, rc)
	CommunityRouter(router, db, rc)
	UserRouter(router, db, rc)
	TestimonyRouter(router, db, rc)
	notificationRouter(router, db, rc)
	OrganizerRouter(router, db, rc)
	AdminRouter(router, db, rc)

	router.GET("api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
