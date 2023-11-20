package messages

import (
	"testing"
)

func TestCantBeBlankMessage(t *testing.T) {
	exp := "can't be blank"

	res := BLANK

	if exp != res {
		t.Errorf(M, exp, res)
	}
}

func TestIsNoValidMessage(t *testing.T) {
	exp := "is not valid"

	res := INVALID

	if exp != res {
		t.Errorf(M, exp, res)
	}
}
