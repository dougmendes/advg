package service

import (
	"context"

	"github.com/dougmendes/advg/pkg/lawyer/model"
)

type service struct {
	repository model.Repository
}

func NewService(r model.Repository) model.Service {
	return &service{
		repository: r,
	}
}

func (s service) GetLawyerById(ctx context.Context, id string) (*model.Lawyer, error) {
	emp, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return emp, nil

}
