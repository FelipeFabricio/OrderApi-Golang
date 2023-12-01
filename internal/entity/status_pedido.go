package entity

type StatusPedido int

const (
	Recebido StatusPedido = iota
	EmPreparacao
	Pronto
	Finalizado
)

var descricaoStatusPedido = [...]string{
	"Recebido",
	"EmPreparacao",
	"Pronto",
	"Finalizado",
}

func (s StatusPedido) ObterDescricaoStatusPedido() string {
	if s < 0 || int(s) >= len(descricaoStatusPedido) {
		return "Desconhecido"
	}
	return descricaoStatusPedido[s]
}
