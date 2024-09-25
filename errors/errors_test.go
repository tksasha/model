package errors_test

import (
	"testing"

	"github.com/tksasha/model/errors"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNew(t *testing.T) {
	t.Run("when no one argument is provided", func(t *testing.T) {
		sbj := errors.New()

		res := sbj.IsEmpty()

		assert.Assert(t, res)
	})

	t.Run("when only one argument is provided", func(t *testing.T) {
		sbj := errors.New("name")

		res := sbj.IsEmpty()

		assert.Assert(t, res)
	})

	t.Run("when two arguments are provided", func(t *testing.T) {
		sbj := errors.New("name", "can't be blank")

		res := sbj.Get("name")

		assert.Assert(t, is.Contains(res, "can't be blank"))
	})
}

func TestSet(t *testing.T) {
	t.Run("when attribute in lowercase", func(t *testing.T) {
		sbj := errors.New()

		sbj.Set("name", "required")

		res := sbj.Get("name")

		assert.Assert(t, is.Contains(res, "required"))
	})

	t.Run("when attribute in uppercase", func(t *testing.T) {
		sbj := errors.New()

		sbj.Set("Name", "required")

		res := sbj.Get("name")

		assert.Assert(t, is.Contains(res, "required"))
	})
}

func TestGet(t *testing.T) {
	sbj := errors.New()

	sbj.Set("name", "can't be blank")

	res := sbj.Get("name")

	assert.Assert(t, is.Contains(res, "can't be blank"))
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		sbj := errors.New()

		res := sbj.IsEmpty()

		assert.Assert(t, res)
	})

	t.Run("when it is not empty", func(t *testing.T) {
		sbj := errors.New()

		sbj.Set("name", "can't be blank")

		res := !sbj.IsEmpty()

		assert.Assert(t, res)
	})
}

func TestCheck(t *testing.T) {
	t.Run("when there are errors for an attribute it should return true", func(t *testing.T) {
		sbj := errors.New("foo", "bar").Check("foo")

		assert.Assert(t, sbj)
	})

	t.Run("when there no errors for an attribute it should return false", func(t *testing.T) {
		sbj := errors.New("foo", "bar").Check("jar")

		assert.Assert(t, !sbj)
	})
}
