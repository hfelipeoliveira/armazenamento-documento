package clientesrv

import (
	"errors"
	"time"

	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/dominio"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/portas"
	"github.com/rs/xid"
)

type servico struct {
	clienteRepositorio portas.ClienteRepositorio
}

func Novo(clienteRepositorio portas.ClienteRepositorio) *servico {
	return &servico{
		clienteRepositorio: clienteRepositorio,
	}
}

func (srv *servico) Recuperar(id string) (*dominio.Cliente, error) {
	cliente, err := srv.clienteRepositorio.RecuperarPorId(id)
	if err != nil {
		return nil, errors.New("Erro ao recuperar cliente do repositório")
	}

	return cliente, nil
}

func (srv *servico) Criar(novoCliente dto.NovoCliente) (*dominio.Cliente, error) {
	cliente := dominio.Cliente{
		ID:          xid.New().String(),
		Cnpj:        novoCliente.Cnpj,
		RazaoSocial: novoCliente.RazaoSocial,
		Situacao:    dominio.CLIENTE_ATIVADO,
		CriadoEm:    time.Now(),
	}

	err := srv.clienteRepositorio.Criar(&cliente)
	if err == nil {
		return nil, err
	}

	return &cliente, nil
}