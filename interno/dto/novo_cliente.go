package dto

type NovoCliente struct {
	Cnpj        string `json:"cnpj" validate:"required,len=14"`
	RazaoSocial string `json:"razao_social" validate:"required,min=3,max=256"`
}
