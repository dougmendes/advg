package config

import (
	. "github.com/dougmendes/advg/pkg/lawyer/controller"
	"github.com/gin-gonic/gin"
)

func ConfigurationRoutes(router *gin.Engine) *gin.Engine {

	baseUrl := router.Group("/api/v1/")
	{
		LawyerRoutes(baseUrl)

	}

	return router
}
