package main

import (
	"encoding/json"
	"os"

	"github.com/tksasha/model"
	"github.com/tksasha/validations"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	model.Model
}

func (user *User) Valid() bool {
	user.Errors = validations.NewErrors()

	validations.ValidatePresenceOf(user.Errors, "name", user.Name)

	return user.Errors.Empty()
}

func NewUser(name string, age int) (user User) {
	user.Name = name

	user.Age = age

	user.Errors = validations.NewErrors()

	return
}

func main() {
	var user User

	user = NewUser("John McClane", 42)

	if user.Valid() {
		json.NewEncoder(os.Stdout).Encode(user)
	} else {
		json.NewEncoder(os.Stdout).Encode(user.Errors)
	}

	user = NewUser("", 42)

	if user.Valid() {
		json.NewEncoder(os.Stdout).Encode(user)
	} else {
		json.NewEncoder(os.Stdout).Encode(user.Errors)
	}
}
