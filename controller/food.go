package controller

import (
	"net/http"

	"github.com/andrasbarabas/shapeshiftr-api/model"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/errs"
	"github.com/andrasbarabas/shapeshiftr-api/service"
	"github.com/gin-gonic/gin"
)

type FoodController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type GORMFoodController struct {
	service service.FoodService
}

func NewFoodController(service service.FoodService) FoodController {
	return &GORMFoodController{
		service: service,
	}
}

func (c *GORMFoodController) Create(ctx *gin.Context) {
	var foodRequest model.FoodRequest

	if err := ctx.ShouldBindJSON(&foodRequest); err != nil {
		errs.HandleError(ctx, []errs.Error{
			*errs.CreateError(errs.GetErrorMessage(errs.BadRequestError), http.StatusBadRequest, nil),
		})

		return
	}

	if err := foodRequest.Validate(); err != nil {
		errs.HandleError(ctx, err)

		return
	}

	result, err := c.service.Create(foodRequest)

	if err != nil {
		errs.HandleError(ctx, []errs.Error{*err})

		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *GORMFoodController) Delete(ctx *gin.Context) {
	refID := model.RefID(ctx.Param("id"))

	if err := refID.Validate(); err != nil {
		errs.HandleError(ctx, err)

		return
	}

	if err := c.service.Delete(refID); err != nil {
		errs.HandleError(ctx, []errs.Error{*err})

		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *GORMFoodController) GetAll(ctx *gin.Context) {
	result, err := c.service.GetAll()

	if err != nil {
		errs.HandleError(ctx, []errs.Error{*err})

		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *GORMFoodController) GetOne(ctx *gin.Context) {
	refID := model.RefID(ctx.Param("id"))

	if err := refID.Validate(); err != nil {
		errs.HandleError(ctx, err)

		return
	}

	result, err := c.service.GetOne(refID)

	if err != nil {
		errs.HandleError(ctx, []errs.Error{*err})

		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *GORMFoodController) Update(ctx *gin.Context) {
	refID := model.RefID(ctx.Param("id"))

	if err := refID.Validate(); err != nil {
		errs.HandleError(ctx, err)

		return
	}

	var foodRequest model.FoodRequest

	if err := ctx.ShouldBindJSON(&foodRequest); err != nil {
		errs.HandleError(ctx, []errs.Error{
			*errs.CreateError(errs.GetErrorMessage(errs.BadRequestError), http.StatusBadRequest, nil),
		})

		return
	}

	if err := foodRequest.Validate(); err != nil {
		errs.HandleError(ctx, err)

		return
	}

	result, err := c.service.Update(refID, foodRequest)

	if err != nil {
		errs.HandleError(ctx, []errs.Error{*err})

		return
	}

	ctx.JSON(http.StatusOK, result)
}
