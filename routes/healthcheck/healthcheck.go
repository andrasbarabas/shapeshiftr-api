package healthcheck

import (
	"github.com/andrasbarabas/shapeshiftr-api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	routerGroup := r.Group("/healthcheck")

	controller := setupController()

	routerGroup.GET("/", controller.Ping)
}

func setupController() controller.HealthcheckController {
	controller := controller.NewHealthcheckController()

	return controller
}
