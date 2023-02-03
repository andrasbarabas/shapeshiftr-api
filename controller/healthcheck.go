package controller

import (
	"net/http"

	"github.com/andrasbarabas/shapeshiftr-api/model"
	"github.com/gin-gonic/gin"
)

type HealthcheckController interface {
	Ping(ctx *gin.Context)
}

type GORMHealthcheckController struct{}

func NewHealthcheckController() HealthcheckController {
	return &GORMHealthcheckController{}
}

func (c *GORMHealthcheckController) Ping(ctx *gin.Context) {
	response := model.Healthcheck{
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, response)
}
