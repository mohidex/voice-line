package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/voice-line/internal/repositories"
	"github.com/mohidex/voice-line/pkg/auth"
)

type UserHandler struct {
	Repo repositories.UserRepository
	Auth auth.Authenticator
}

func (uh UserHandler) SignUP(c *gin.Context) {}

func (uh UserHandler) SignIn(c *gin.Context) {}

func (uh UserHandler) UserInfo(c *gin.Context) {}
