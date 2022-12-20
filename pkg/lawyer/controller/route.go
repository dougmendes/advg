package controller

import (
	"github.com/dougmendes/advg/connections"
	"github.com/dougmendes/advg/pkg/lawyer/repository"
	"github.com/dougmendes/advg/pkg/lawyer/service"
	"github.com/gin-gonic/gin"
)

func LawyerRoutes(group *gin.RouterGroup) {
	lawyerGroup := group.Group("/lawyer")
	{
		lawyerRepo := repository.NewRepository(connections.NewConnection())
		lawyerService := service.NewService(lawyerRepo)
		l := NewLawyerController(lawyerService)

		lawyerGroup.GET("/:id", l.GetLawyerById())
	}
}
