package rest

import (
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/portas"
)

type Manipulador struct {
	ClienteServico portas.ClienteServico
}

func NovoManipulador(clienteServico portas.ClienteServico) Manipulador {
	return Manipulador{
		ClienteServico: clienteServico,
	}
}
