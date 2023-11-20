package model

import (
	"testing"

	"github.com/tksasha/validations"
)

const M = "\033[31m`%v` was expected, but it is `%v`\033[0m"

func TestIsValid(t *testing.T) {
	model := new(Model)

	t.Run("when it is valid", func(t *testing.T) {
		model.Errors = validations.NewErrors()

		exp := true

		res := model.IsValid()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is not valid", func(t *testing.T) {
		model.Errors = validations.NewErrors()

		model.Errors.Add("attribute", "can't be blank")

		exp := false

		res := model.IsValid()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}

func TestIsNotValid(t *testing.T) {
	model := new(Model)

	t.Run("when it is valid", func(t *testing.T) {
		model.Errors = validations.NewErrors()

		exp := false

		res := model.IsNotValid()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is not valid", func(t *testing.T) {
		model.Errors = validations.NewErrors()

		model.Errors.Add("attribute", "is not valid")

		exp := true

		res := model.IsNotValid()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}
