package model

import (
	"github.com/tksasha/model/errors"
)

type Model struct {
	Errors errors.ValidationError `json:"errors"`
}

func New() *Model {
	return &Model{
		Errors: errors.New(),
	}
}

func (model *Model) IsValid() bool {
	return model.Errors.IsEmpty()
}
