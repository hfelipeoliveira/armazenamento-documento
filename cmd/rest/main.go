package main

import (
	"log"
	"os"

	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/servicos/clientesrv"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/repositorios"
	"github.com/joho/godotenv"
)

func main() {
	lerVariaveisDeAmbiente()
	conexaoConfig := dto.ConexaoMysqlConfig{
		Usuario:     os.Getenv("DB_USUARIO"),
		Senha:       os.Getenv("DB_SENHA"),
		Host:        os.Getenv("DB_HOST"),
		Porta:       os.Getenv("DB_PORTA"),
		BaseDeDados: os.Getenv("DB_NOME"),
	}
	mysql, err := repositorios.NovaConexaoMysql(conexaoConfig)

	if err != nil {
		panic(err)
	}

	clienteRepositorio := repositorios.ClienteRepositorio{Db: mysql}

	clienteServico := clientesrv.Novo(&clienteRepositorio)
	clienteServico.Criar(dto.NovoCliente{
		Cnpj:        "12345678901234",
		RazaoSocial: "Teste massa",
	})

}

func lerVariaveisDeAmbiente() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Erro ao ler arquivo .env")
	}
}
