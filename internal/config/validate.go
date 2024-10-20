package config

import "github.com/go-playground/validator/v10"

func Valid(struc interface{}) error {
	v := validator.New()

	if err := v.Struct(struc); err != nil {
		return err
	}

	return nil
}