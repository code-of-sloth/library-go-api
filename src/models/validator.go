package models

import (
	"errors"
	"fmt"
)

type Validatable interface {
	Validate() error
}

var ErrNotValidatable = errors.New("type is not validatable")

type Validator struct{}

func (v *Validator) Validate(i any) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	return fmt.Errorf("bindings:internal error during validation")
}
