package validadorsrv

import "github.com/go-playground/validator"

func Novo() Servico {
	return Servico{
		validator: validator.New(),
	}
}

type Servico struct {
	validator *validator.Validate
}

func (s *Servico) Validar(objeto interface{}) error {
	// Parei aqui
	return nil
}
