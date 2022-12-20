package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	. "github.com/dougmendes/advg/pkg/errs"
	"github.com/dougmendes/advg/pkg/lawyer/model"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetById(ctx context.Context, id int) (*model.Lawyer, *Error) {
	row := r.db.QueryRow(queryGetById, id)

	var l model.Lawyer
	err := row.Scan(&l.Id, &l.Name, &l.LastName, &l.Resume, &l.Oab, &l.Tel, &l.ProfilePicture)
	if err == sql.ErrNoRows {
		return nil, ErrorResponse(http.StatusNotFound, fmt.Sprintf("lawyer not found for id %v", id), err)
	}
	if err != nil {
		return nil, ErrorResponse(http.StatusInternalServerError, "Unexpected error", err)
	}

	return &l, nil
}

func NewRepository(db *sql.DB) model.Repository {
	return &repository{
		db: db,
	}
}
