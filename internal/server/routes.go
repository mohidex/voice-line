package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/voice-line/internal/handlers"
	"github.com/mohidex/voice-line/internal/middlewares"
	"github.com/mohidex/voice-line/internal/repositories"
	"github.com/mohidex/voice-line/pkg/auth"
)

type Routes struct {
	UserRepo repositories.UserRepository
	Auth     auth.Authenticator
}

func (r *Routes) Setup(router *gin.Engine) {

	health := &handlers.HealthCheckHandler{}

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userHandler := &handlers.UserHandler{
			Repo: r.UserRepo,
			Auth: r.Auth,
		}
		v1.POST("/signup", userHandler.SignUP)
		v1.POST("/login", userHandler.SignIn)

		userRoutes := v1.Group("/users")
		userRoutes.Use(middlewares.AuthMiddleware(r.Auth))
		userRoutes.GET("/me", userHandler.UserInfo)
	}
}
