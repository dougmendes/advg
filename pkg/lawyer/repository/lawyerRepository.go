package repository

import (
	"context"

	"github.com/dougmendes/advg/pkg/lawyer/model"
)

type repository struct {
}

func (r *repository) GetById(ctx context.Context, id string) (*model.Lawyer, error) {
	//lawyers := make([]model.LawyerResponse, 0)
	douglas := model.Lawyer{Id: "jqjhwwGg"}
	//lawyers = append(lawyers, douglas)
	return &douglas, nil
}

func NewRepository() model.Repository {
	return &repository{}
}
