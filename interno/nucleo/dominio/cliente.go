package dominio

import (
	"time"
)

const (
	CLIENTE_ATIVADO = "ativado"
)

type Cliente struct {
	ID          string     `gorm:"primaryKey;size:50;not null"`
	Cnpj        string     `gorm:"size:14;not null"`
	RazaoSocial string     `gorm:"size:256;not null"`
	Situacao    string     `gorm:"size:50;not null"`
	CriadoEm    time.Time  `gorm:"type:datetime;not null"`
	EditadoEm   *time.Time `gorm:"type:datetime;null"`
}
