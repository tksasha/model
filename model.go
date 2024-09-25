package model

import (
	"github.com/tksasha/model/errors"
)

type Model struct {
	Errors errors.Errors `json:"errors"`
}

func New() *Model {
	return &Model{
		Errors: errors.New(),
	}
}

func (model *Model) Valid() bool {
	return model.Errors.Empty()
}
