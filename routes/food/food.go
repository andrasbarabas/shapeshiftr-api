package food

import (
	"github.com/andrasbarabas/shapeshiftr-api/controller"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/db"
	"github.com/andrasbarabas/shapeshiftr-api/repository"
	"github.com/andrasbarabas/shapeshiftr-api/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	routerGroup := r.Group("/foods")

	controller := setupController()

	routerGroup.GET("/", controller.GetAll)
	routerGroup.GET("/:id", controller.GetOne)
	routerGroup.POST("/", controller.Create)
	routerGroup.PUT("/:id", controller.Update)
	routerGroup.DELETE("/:id", controller.Delete)
}

func setupController() controller.FoodController {
	service := setupService()

	controller := controller.NewFoodController(service)

	return controller
}

func setupRepository() repository.FoodRepository {
	repository := repository.NewFoodRepository(db.GetConnection())

	return repository
}

func setupService() service.FoodService {
	repository := setupRepository()

	service := service.NewFoodService(repository)

	return service
}
