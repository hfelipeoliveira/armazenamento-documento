package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/hfelipeoliveira/armazenamento-documento/interno/dto"
	"github.com/mitchellh/mapstructure"
)

func (m *Manipulador) ClienteRecuperarPorId(w http.ResponseWriter, r *http.Request) {
	clinteId := chi.URLParam(r, "id")
	cliente, err := m.ClienteServico.Recuperar(clinteId)

	if err != nil {
		render.JSON(w, r, map[string]string{"erro": err.Error()})
		return
	}

	var clienteContrato dto.ClienteResposta

	err = mapstructure.Decode(cliente, &clienteContrato)

	if err != nil {
		render.JSON(w, r, map[string]string{"erro": err.Error()})
		return
	}

	render.JSON(w, r, clienteContrato)
}
