package model

import (
	"testing"

	"github.com/tksasha/validations"
)

func TestIsValid(t *testing.T) {
	model := Model{}

	t.Run("when it is valid", func(t *testing.T) {
		model.Errors = validations.NewErrors()

		expected := true

		result := model.IsValid()

		if expected != result {
			t.Errorf(validations.M, expected, result)
		}
	})

	t.Run("when it is not valid", func(t *testing.T) {
		model.Errors = validations.NewErrors()

		model.Errors.Add("attribute", "can't be blank")

		expected := false

		result := model.IsValid()

		if expected != result {
			t.Errorf(validations.M, expected, result)
		}
	})
}
