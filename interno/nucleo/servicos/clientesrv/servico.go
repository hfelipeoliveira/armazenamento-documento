package clientesrv

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/dominio"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/portas"
	"github.com/rs/xid"
)

type Servico struct {
	clienteRepositorio portas.ClienteRepositorio
	validate           *validator.Validate
}

func Novo(clienteRepositorio portas.ClienteRepositorio) *Servico {
	return &Servico{
		clienteRepositorio: clienteRepositorio,
		validate:           validator.New(),
	}
}

func (srv *Servico) RecuperarPorId(id string) (*dominio.Cliente, error) {
	cliente, err := srv.clienteRepositorio.RecuperarPorId(id)
	if err != nil {
		return nil, errors.New("Erro ao recuperar cliente do repositório")
	}

	return cliente, nil
}

func (srv *Servico) Listar() ([]*dominio.Cliente, error) {
	cliente, err := srv.clienteRepositorio.Listar()
	if err != nil {
		return nil, errors.New("Erro ao recuperar clientes do repositório")
	}

	return cliente, nil
}

func (srv *Servico) Criar(novoCliente dto.NovoCliente) (*dominio.Cliente, error) {
	err := srv.validate.Struct(&novoCliente)
	if err != nil {
		return nil, err
	}

	cliente := dominio.Cliente{
		ID:          xid.New().String(),
		Cnpj:        novoCliente.Cnpj,
		RazaoSocial: novoCliente.RazaoSocial,
		Situacao:    dominio.CLIENTE_ATIVADO,
		CriadoEm:    time.Now(),
	}

	err = srv.clienteRepositorio.Criar(&cliente)
	if err != nil {
		return nil, err
	}

	return &cliente, nil
}
