package model

import (
	"testing"
)

func TestAdd(t *testing.T) {
	exp := "can't be blank"

	errs := NewErrors()

	errs.Add("name", "can't be blank")

	res := errs.Get("name")[0]

	if exp != res {
		t.Errorf(M, exp, res)
	}
}

func TestGet(t *testing.T) {
	exp := "can't be blank"

	errs := NewErrors()

	errs.Add("name", "can't be blank")

	res := errs.Get("name")[0]

	if exp != res {
		t.Errorf(M, exp, res)
	}
}

func TestIsEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		exp := true

		errs := NewErrors()

		res := errs.IsEmpty()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is not empty", func(t *testing.T) {
		exp := false

		errs := NewErrors()

		errs.Add("name", "can't be blank")

		res := errs.IsEmpty()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("when it is empty", func(t *testing.T) {
		exp := false

		errs := NewErrors()

		res := errs.IsNotEmpty()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is not empty", func(t *testing.T) {
		exp := true

		errs := NewErrors()

		errs.Add("name", "can't be blank")

		res := errs.IsNotEmpty()

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}
