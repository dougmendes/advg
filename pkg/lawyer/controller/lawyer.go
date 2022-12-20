package controller

import (
	"net/http"
	"strconv"

	"github.com/dougmendes/advg/pkg/lawyer/model"
	"github.com/dougmendes/advg/pkg/response"
	"github.com/gin-gonic/gin"
)

func NewLawyerController(s model.Service) *LawyerController {
	return &LawyerController{
		service: s,
	}
}

type LawyerController struct {
	service model.Service
}

func (lwy *LawyerController) GetLawyerById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, _ := strconv.ParseInt(idStr, 10, 0)
		lawyer, err := lwy.service.GetLawyerById(ctx, int(id))
		if err != nil {
			ctx.JSON(err.Code, response.DecodeError(err))
		}
		ctx.JSON(http.StatusOK, response.LaweryResponse(lawyer))
	}
}
