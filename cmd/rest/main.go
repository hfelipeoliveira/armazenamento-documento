package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/manipuladores/rest"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/servicos/clientesrv"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/nucleo/servicos/validadorsrv"
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

	validador := validadorsrv.Novo()

	clienteRepositorio := repositorios.ClienteRepositorio{Db: mysql}

	clienteServico := clientesrv.Novo(&clienteRepositorio, &validador)
	clienteServico.Criar(dto.NovoCliente{
		Cnpj:        "123456789012345",
		RazaoSocial: "Teste massa",
	})

	manipulador := rest.NovoManipulador(clienteServico)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/rest/v1/clientes", func(r chi.Router) {
		r.Get("/", manipulador.ClienteListar)
		r.Post("/", manipulador.ClienteCriar)
		r.Get("/{id}", manipulador.ClienteRecuperarPorId)
	})

	http.ListenAndServe(":80", r)
}

func lerVariaveisDeAmbiente() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Erro ao ler arquivo .env")
	}
}
