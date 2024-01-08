package dominio

import (
	"database/sql"
	"time"
)

const (
	CLIENTE_ATIVADO = "ativado"
)

type Cliente struct {
	ID          string    `gorm:"size:50;not null"`
	Cnpj        string    `gorm:"size:14;not null"`
	RazaoSocial string    `gorm:"size:256;not null"`
	Situacao    string    `gorm:"size:50;not null"`
	CriadoEm    time.Time `gorm:"not null"`
	EditadoEm   sql.NullTime
}
