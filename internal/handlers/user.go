package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/voice-line/internal/models"
	"github.com/mohidex/voice-line/internal/repositories"
	"github.com/mohidex/voice-line/pkg/auth"
)

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

type UserHandler struct {
	Repo repositories.UserRepository
	Auth auth.Authenticator
}

func (uh UserHandler) SignUP(c *gin.Context) {
	var input registerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := uh.Auth.CreateUser(c, input.Email, input.Password1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.NewUser(response.LocalID, input.Name, response.Email, true, false)

	err = uh.Repo.CreateUser(c, newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fetch user from database
	user, err := uh.Repo.GetUserByID(c, response.LocalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})

}

func (uh UserHandler) SignIn(c *gin.Context) {
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := uh.Auth.GetToken(c, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (uh UserHandler) UserInfo(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	user, err := uh.Repo.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
