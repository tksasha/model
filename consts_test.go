package model

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestErrIsBlank(t *testing.T) {
	assert.Equal(t, ErrIsBlank, "can't be blank")
}

func TestErrIsNoValid(t *testing.T) {
	assert.Equal(t, ErrIsNotValid, "is not valid")
}

func TestErrExecSQL(t *testing.T) {
	assert.Equal(t, ErrExecSQL, "couldn't exec SQL")
}

func TestErrObtainID(t *testing.T) {
	assert.Equal(t, ErrObtainID, "couldn't obtain ID")
}
