package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func (h HealthCheckHandler) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
