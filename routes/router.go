package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/andrasbarabas/shapeshiftr-api/pkg/errs"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/l"
	"github.com/andrasbarabas/shapeshiftr-api/routes/food"
	"github.com/andrasbarabas/shapeshiftr-api/routes/healthcheck"
	ginzap "github.com/gin-contrib/zap"

	"github.com/andrasbarabas/shapeshiftr-api/config"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := initRouter()

	// Setup router group
	routes := router.Group("/api/")

	healthcheck.SetupRoutes(routes)
	food.SetupRoutes(routes)

	// Setup trusted proxies
	setupTrustedProxies(router)

	return router
}

func setupTrustedProxies(r *gin.Engine) {
	trustedProxies := strings.Split(config.ServerConfig.TrustedProxies, ",")

	err := r.SetTrustedProxies(trustedProxies)

	if err != nil {
		l.Logger.Error(err.Error())
	}
}

func initRouter() *gin.Engine {
	router := gin.New()

	router.Use(ginzap.Ginzap(l.Logger, time.RFC3339, true), gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		err := errs.CreateError(
			errs.GetErrorMessage(errs.ResourceNotFoundError),
			http.StatusNotFound,
			nil,
		)

		errs.HandleError(c, []errs.Error{*err})
	})

	return router
}
