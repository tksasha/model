package errors

import (
	"github.com/tksasha/utils/strings"
)

type Errors interface {
	Set(attribute, message string)
	Get(attribute string) []string
	IsEmpty() bool
}

type errors map[string][]string

func New(args ...string) Errors {
	errs := errors{}

	if len(args) == 2 { //nolint:mnd
		errs.Set(args[0], args[1])
	}

	return errs
}

func (e errors) Set(attribute, message string) {
	attribute = strings.ToSnakeCase(attribute)

	e[attribute] = append(e[attribute], message)
}

func (e errors) Get(attribute string) []string {
	errs, ok := e[attribute]
	if !ok {
		return []string{}
	}

	return errs
}

func (e errors) IsEmpty() bool {
	return e.size() == 0
}

func (e errors) size() int {
	return len(e)
}
