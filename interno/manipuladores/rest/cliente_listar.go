package rest

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/mitchellh/mapstructure"
)

func (m *Manipulador) ClienteListar(w http.ResponseWriter, r *http.Request) {
	clientes, err := m.ClienteServico.Listar()

	if err != nil {
		render.JSON(w, r, map[string]string{"erro": err.Error()})
		return
	}

	var clienteContrato []dto.ClienteResposta

	err = mapstructure.Decode(clientes, &clienteContrato)

	if err != nil {
		render.JSON(w, r, map[string]string{"erro": err.Error()})
		return
	}

	render.JSON(w, r, clienteContrato)
}
