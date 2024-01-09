package dto

type ClienteResposta struct {
	ID          string `json:"id"`
	Cnpj        string `json:"cnpj"`
	RazaoSocial string `json:"razao_social"`
	Situacao    string `json:"situacao"`
}
