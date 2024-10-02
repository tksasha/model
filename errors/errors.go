package errors

import (
	"github.com/tksasha/utils/strings"
)

type Errors map[string][]string

func New(args ...string) Errors {
	errs := Errors{}

	if len(args) == 2 { //nolint:mnd
		errs.Set(args[0], args[1])
	}

	return errs
}

func (e Errors) Set(attribute, message string) {
	attribute = strings.ToSnakeCase(attribute)

	e[attribute] = append(e[attribute], message)
}

func (e Errors) Get(attribute string) []string {
	errs, ok := e[attribute]
	if !ok {
		return []string{}
	}

	return errs
}

func (e Errors) IsEmpty() bool {
	return len(e) == 0
}

func (e Errors) Has(attribute string) bool {
	return len(e[attribute]) > 0
}
