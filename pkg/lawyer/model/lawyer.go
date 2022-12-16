package model

import "context"

type Repository interface {
	GetById(ctx context.Context, id string) (*Lawyer, error)
}

type Service interface {
	GetLawyerById(ctx context.Context, id string) (*Lawyer, error)
}

type Lawyer struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	Resume         string `json:"resume"`
	Oab            string `json:"oab"`
	Tel            string `json:"tel"`
	ProfilePicture string `json:"profile_picture"`
	Password       string `json:"password"`
}
