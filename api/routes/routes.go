package routes

import (
	"go.uber.org/fx"
	// "github.com/gin-gonic/gin"
)


// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewCommonRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// var router = gin.Default()


// func getRoutes() {
// 	v1 := router.Group("/v1")
// 	addPingRoutes(v1)

// 	v2 := router.Group("/v2")
// 	addPingRoutes(v2)
// }

// NewRoutes sets up routes
func NewRoutes(
	commonRoutes CommonRoutes,
	userRoutes UserRoutes,
	authRoutes AuthRoutes,
) Routes {
	return Routes{
		commonRoutes,
		userRoutes,
		authRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}