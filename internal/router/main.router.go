package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MainRouter(router *gin.Engine, db *pgxpool.Pool) {
	authRouter(router, db)
	EventRouter(router, db)
	CommunityRouter(router, db)
	UserRouter(router, db)
	TestimonyRouter(router, db)
	notificationRouter(router, db)
	OrganizerRouter(router, db)
	AdminRouter(router, db)
}
