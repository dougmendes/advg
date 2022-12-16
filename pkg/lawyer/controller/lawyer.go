package controller

import (
	"net/http"

	"github.com/dougmendes/advg/pkg/lawyer/model"
	"github.com/dougmendes/advg/pkg/lawyer/repository"
	"github.com/dougmendes/advg/pkg/lawyer/service"
	"github.com/dougmendes/advg/pkg/response"
	"github.com/gin-gonic/gin"
)

func NewController(s model.Service) *LawyerController {
	return &LawyerController{
		service: s,
	}
}

type LawyerController struct {
	service model.Service
}

func LawyerRoutes(group *gin.RouterGroup) {
	lawyerGroup := group.Group("/lawyer")
	{
		lawyerRepo := repository.NewRepository()
		lawyerService := service.NewService(lawyerRepo)
		l := NewController(lawyerService)

		lawyerGroup.GET("/:id", l.getLawyerById())
	}
}

func (lwy *LawyerController) getLawyerById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		lawyer, err := lwy.service.GetLawyerById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{

				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, response.NewResponse(lawyer))
	}
}
