package repositorios

import (
	"fmt"

	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/dominio"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NovaConexaoMysql(config dto.ConexaoMysqlConfig) (*gorm.DB, error) {
	dsn := criarDSN(config)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		db.AutoMigrate(&dominio.Cliente{})
	}

	return db, err
}

func criarDSN(config dto.ConexaoMysqlConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Usuario,
		config.Senha,
		config.Host,
		config.Porta,
		config.BaseDeDados,
	)
}
