package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type lawyer struct {
	name           string `json:"name"`
	lastName       string `json:"last_name"`
	oab            string `json:"oab"`
	tel            string `json:"tel"`
	profilePicture string `json:"profile_picture"`
	password       string `json:"password"`
}

func NewLawyer() *LawyerController {
	return &LawyerController{}
}

type LawyerController struct {
}

func LawyerRoutes(group *gin.RouterGroup) {
	lawyerRouterGroup := group.Group("/")
	lwy := NewLawyer()
	{
		lawyerRouterGroup.GET("lawyer", lwy.GetLawyer())
	}
}

func (lwy *LawyerController) GetLawyer() gin.HandlerFunc {
	lawyer := lawyer{
		name:           "Douglas",
		lastName:       "Mendes",
		oab:            "1111111",
		tel:            "(31)99999999",
		profilePicture: "http://adsadasdsa.com",
		password:       "Xsdffqeqw",
	}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":            lawyer.name,
			"last_name":       lawyer.lastName,
			"oab":             lawyer.oab,
			"tel":             lawyer.tel,
			"profile_picture": lawyer.profilePicture,
			"password":        lawyer.password,
		})
	}
}
