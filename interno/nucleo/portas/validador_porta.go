package portas

type Validador interface {
	Validar(interface{}) error
}
