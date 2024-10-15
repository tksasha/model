package errors_test

import (
	"testing"

	"github.com/tksasha/model/errors"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNew(t *testing.T) {
	err := errors.New()

	assert.Assert(t, err.IsEmpty())
}

func TestNewWithError(t *testing.T) {
	err := errors.NewWithError("name", "can't be blank")

	assert.Assert(t, err.Has("name"))
	assert.Assert(t, is.Contains(err.Get("name"), "can't be blank"))
}

func TestValidationError_String(t *testing.T) {
	err := errors.NewWithError("name", "can't be blank")

	assert.Equal(t, err.String(), "name: can't be blank")
}

func TestSet(t *testing.T) {
	t.Run("when attribute in lowercase", func(t *testing.T) {
		err := errors.New()

		err.Set("name", "is required")

		assert.Assert(t, is.Contains(err.Get("name"), "is required"))
	})

	t.Run("when attribute in uppercase", func(t *testing.T) {
		err := errors.New()

		err.Set("Name", "is required")

		assert.Assert(t, is.Contains(err.Get("name"), "is required"))
	})
}

func TestGet(t *testing.T) {
	t.Run("when errors are empty it should return an empty array", func(t *testing.T) {
		err := errors.New()

		assert.Assert(t, len(err.Get("name")) == 0)
	})

	t.Run("when errors are not empty it should return a slice with that error", func(t *testing.T) {
		err := errors.NewWithError("name", "can't be blank")

		assert.Assert(t, is.Contains(err.Get("name"), "can't be blank"))
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("when there are no errors it should return true", func(t *testing.T) {
		err := errors.New()

		assert.Assert(t, err.IsEmpty())
	})

	t.Run("when there are errors it should return false", func(t *testing.T) {
		err := errors.NewWithError("name", "can't be blank")

		assert.Assert(t, !err.IsEmpty())
	})
}

func TestHas(t *testing.T) {
	err := errors.NewWithError("foo", "bar")

	t.Run("when there are errors for an attribute it should return true", func(t *testing.T) {
		assert.Assert(t, err.Has("foo"))
	})

	t.Run("when there no errors for an attribute it should return false", func(t *testing.T) {
		assert.Assert(t, !err.Has("jar"))
	})
}
