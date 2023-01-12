package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhhuynh/gin-base/lib"
)

// CommonController data type
type CommonController struct {
	logger lib.Logger
}

// NewCommonController creates new common controller
func NewCommonController(logger lib.Logger) CommonController {
	return CommonController{
		logger: logger,
	}
}

// Ping
func (common CommonController) Ping(c *gin.Context) {
	message := "pong"
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}

// HealthCheck
func (common CommonController) HealthCheck(c *gin.Context) {
	message := "Welcome to Gin Base"
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}
