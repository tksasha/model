package model_test

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/tksasha/model"
	"gotest.tools/v3/assert"
)

func TestIsValid(t *testing.T) {
	sbj := model.New()

	t.Run("when there are no errors it should return true", func(t *testing.T) {
		res := sbj.IsValid()

		assert.Assert(t, res)
	})

	t.Run("when there are errors it should return false", func(t *testing.T) {
		sbj.Errors.Set("attribute", "can't be blank")

		res := sbj.IsValid()

		assert.Assert(t, !res)
	})
}

func TestMarshalJSON(t *testing.T) {
	sbj := model.New()

	sbj.Errors.Set("attribute", "can't be blank")

	bts, err := json.Marshal(sbj)

	exp := []byte(`{"errors":{"attribute":["can't be blank"]}}`)

	res := slices.Equal(bts, exp)

	assert.NilError(t, err)
	assert.Assert(t, res)
}
