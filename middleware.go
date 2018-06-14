package main

import "github.com/go-playground/validator"

/**
 * Enable validator in global context
 */
type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func init() {
	// validator
	e.Validator = &Validator{validator: validator.New()}
}
