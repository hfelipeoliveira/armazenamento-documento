package repositorios

import (
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/dominio"
	"gorm.io/gorm"
)

type ClienteRepositorio struct {
	Db *gorm.DB
}

func (c *ClienteRepositorio) RecuperarPorId(id string) (*dominio.Cliente, error) {
	var cliente dominio.Cliente
	tx := c.Db.Preload("clientes").First(&cliente, "id = ?", id)
	return &cliente, tx.Error
}

func (c *ClienteRepositorio) Criar(cliente *dominio.Cliente) error {
	tx := c.Db.Create(cliente)
	return tx.Error
}
