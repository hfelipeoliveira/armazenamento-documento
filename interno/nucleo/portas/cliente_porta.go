package portas

import (
	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/dominio"
)

type ClienteRepositorio interface {
	RecuperarPorId(id string) (*dominio.Cliente, error)
	Listar() ([]*dominio.Cliente, error)
	Criar(*dominio.Cliente) error
}

type ClienteServico interface {
	RecuperarPorId(id string) (*dominio.Cliente, error)
	Listar() ([]*dominio.Cliente, error)
	Criar(dto.NovoCliente) (*dominio.Cliente, error)
}
