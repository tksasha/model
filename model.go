package model

import (
	"github.com/tksasha/validations"
)

type Model struct {
	Errors validations.Errors `json:"-"`
}

func (model *Model) IsValid() bool {
	return model.Errors.IsEmpty()
}
