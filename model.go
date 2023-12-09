package model

import (
	"github.com/tksasha/model/errors"
)

type Model struct {
	Errors errors.Errors `json:"-"`
}

func (model *Model) IsValid() bool {
	return model.Errors.IsEmpty()
}
