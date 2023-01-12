package bootstrap

import (
	"github.com/thinhhuynh/gin-base/api/controllers"
	"github.com/thinhhuynh/gin-base/api/middlewares"
	"github.com/thinhhuynh/gin-base/api/routes"
	"github.com/thinhhuynh/gin-base/lib"
	"github.com/thinhhuynh/gin-base/repository"
	"github.com/thinhhuynh/gin-base/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
