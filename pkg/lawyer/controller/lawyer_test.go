package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/dougmendes/advg/pkg/lawyer/controller"
	"github.com/dougmendes/advg/pkg/lawyer/model"
	mock_model "github.com/dougmendes/advg/pkg/lawyer/model/mock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

const (
	relativePathLawyers = "/api/v1/lawyer/"
	target              = "api/v1/lawyer/:id"
)

func callMock(t *testing.T) (*mock_model.MockService, *LawyerController) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_model.NewMockService(ctrl)
	handler := NewLawyerController(service)

	return service, handler
}

func TestController_ById(t *testing.T) {
	lwy := &model.Lawyer{
		Id:             1,
		Name:           "Douglas",
		LastName:       "Mendes",
		Resume:         "I'm Douglas",
		Oab:            "123124",
		Tel:            "(31)991399147",
		ProfilePicture: "https://picture.com.br",
	}

	service, handler := callMock(t)

	api := gin.New()
	api.GET(target, handler.GetLawyerById())
	service.EXPECT().GetLawyerById(gomock.Any(), int(1)).Return(lwy, nil)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/lawyer/1", nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestController_ById_NotFound(t *testing.T) {
	service, handler := callMock(t)
	api := gin.New()
	api.GET(target, handler.GetLawyerById())
	service.EXPECT().GetLawyerById(gomock.Any(), int(3)).Return(nil, errors.New("Lawyer don't finded for id 3"))
	req := httptest.NewRequest(http.MethodGet, "/api/v1/lawyer/3", nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.NotEqual(t, http.StatusOK, resp.Code)
}
