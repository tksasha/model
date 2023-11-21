package model

import (
	"database/sql"
)

type Model struct {
	DB     *sql.DB `json:"-"`
	Errors Errors  `json:"-"`
}

func (model *Model) IsValid() bool {
	return model.Errors.IsEmpty()
}

func (model *Model) IsNotValid() bool {
	return model.Errors.IsNotEmpty()
}
