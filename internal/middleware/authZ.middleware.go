package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MGT06/EventHub_Backend.git/internal/dto"
	"github.com/MGT06/EventHub_Backend.git/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func CheckToken(rc *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "please login first",
			})
			return
		}

		result := strings.Split(bearer, " ")
		if len(result) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}
		if result[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}

		var token pkg.JWTClaims
		err := token.DecodeToken(result[1])
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
					Success: false,
					Message: "invalid token",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "a system error has occurred",
			})
			return
		}

		key := fmt.Sprintf("eventhub:tokenBlacklist:%s", token.ID)
		n, err := rc.Exists(c.Request.Context(), key).Result()
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, dto.Response{
				Success: false,
				Message: "service unavailable",
			})
			return
		}
		if n > 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "invalid token",
			})
			return
		}

		c.Set("idUser", token.Id)
		c.Set("role", token.Role)
		c.Set("jti", token.ID)
		c.Set("exp", token.ExpiresAt.Time)
		c.Next()
	}
}

func OrganizerAccess(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "login first",
		})
		return
	}

	if role != "organizer" {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "no privilege",
		})
		return
	}
}

func AdminAccess(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "login first",
		})
		return
	}

	if role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "no privilege",
		})
		return
	}
}
