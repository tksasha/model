package errors

import (
	"fmt"
	"strings"

	utilsstrings "github.com/tksasha/utils/strings"
)

type ValidationError map[string][]string

func New() ValidationError {
	return ValidationError{}
}

func NewWithError(attribute, message string) ValidationError {
	err := ValidationError{}

	err.Set(attribute, message)

	return err
}

func (e ValidationError) String() string {
	messages := []string{}

	for attribute, errs := range e {
		messages = append(messages, fmt.Sprintf("%s: %s", attribute, strings.Join(errs, ", ")))
	}

	return strings.Join(messages, "; ")
}

func (e ValidationError) Set(attribute, message string) {
	attribute = utilsstrings.ToSnakeCase(attribute)

	e[attribute] = append(e[attribute], message)
}

func (e ValidationError) Get(attribute string) []string {
	return e[attribute]
}

func (e ValidationError) IsEmpty() bool {
	return len(e) == 0
}

func (e ValidationError) Has(attribute string) bool {
	return len(e[attribute]) > 0
}
