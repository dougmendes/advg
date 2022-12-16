package controller

import (
	"net/http"
	"strconv"

	"github.com/dougmendes/advg/connections"
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
		lawyerRepo := repository.NewRepository(connections.NewConnection())
		lawyerService := service.NewService(lawyerRepo)
		l := NewController(lawyerService)

		lawyerGroup.GET("/:id", l.getLawyerById())
	}
}

func (lwy *LawyerController) getLawyerById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, _ := strconv.ParseInt(idStr, 10, 0)
		lawyer, err := lwy.service.GetLawyerById(ctx, int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{

				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, response.NewResponse(lawyer))
	}
}
