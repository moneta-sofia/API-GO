package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("first name required")
var ErrLastNameRequired = errors.New("last name required")

type ErrorNotFound struct {
	ID uint64
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("user not found with ID %d", e.ID)
}
