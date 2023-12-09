package model_test

import (
	"testing"

	"github.com/tksasha/model"
	"github.com/tksasha/model/errors"
	"gotest.tools/v3/assert"
)

var (
	subject = new(model.Model)
)

func TestIsValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = errors.New()

		assert.Assert(t, subject.IsValid())
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = errors.New()

		subject.Errors.Add("attribute", "can't be blank")

		assert.Assert(t, !subject.IsValid())
	})
}

func TestIsNotValid(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		subject.Errors = errors.New()

		assert.Assert(t, subject.IsValid())
	})

	t.Run("when it is not valid", func(t *testing.T) {
		subject.Errors = errors.New()

		subject.Errors.Add("attribute", "is not valid")

		assert.Assert(t, !subject.IsValid())
	})
}
