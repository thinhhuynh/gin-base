package routes

import (
	"github.com/thinhhuynh/gin-base/api/controllers"
	"github.com/thinhhuynh/gin-base/api/middlewares"
	"github.com/thinhhuynh/gin-base/lib"
)

// CommonRoutes struct
type CommonRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	commonController controllers.CommonController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup common routes
func (s CommonRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api").Use()
	{
		api.GET("/ping", s.commonController.Ping)
		api.GET("/healthcheck", s.commonController.HealthCheck)
	}
}

// NewCommonRoutes creates new user controller
func NewCommonRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	commonController controllers.CommonController,
	authMiddleware middlewares.JWTAuthMiddleware,
) CommonRoutes {
	return CommonRoutes{
		handler:        handler,
		logger:         logger,
		commonController: commonController,
		authMiddleware: authMiddleware,
	}
}
