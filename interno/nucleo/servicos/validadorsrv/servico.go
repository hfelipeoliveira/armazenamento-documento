package validadorsrv

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

func Novo() Validador {
	return Validador{
		validator: validator.New(),
	}
}

type Validador struct {
	validator *validator.Validate
}

func (s *Validador) Validar(objeto interface{}) error {
	err := s.validator.Struct(objeto)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " é requerido")
	case "len":
		return errors.New(field + " deve ter " + validationError.Param() + " caracteres")
	case "max":
		return errors.New(field + " é requerido com no máximo " + validationError.Param())
	case "min":
		return errors.New(field + " é requerido com no mínimo " + validationError.Param())
	case "email":
		return errors.New(field + " é inválido")
	}
	return nil
}
