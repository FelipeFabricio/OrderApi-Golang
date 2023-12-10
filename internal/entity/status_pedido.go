package entity

type StatusPedido int

const (
	AguardandoPagamento StatusPedido = iota
	Recebido
	EmPreparacao
	Pronto
	Finalizado
	Cancelado
)

var descricaoStatusPedido = [...]string{
	"Aguardando pagamento",
	"Recebido",
	"Em preparação",
	"Pronto",
	"Finalizado",
	"Cancelado",
}

func (s StatusPedido) ObterDescricaoStatusPedido() string {
	if s < 0 || int(s) >= len(descricaoStatusPedido) {
		return "Desconhecido"
	}
	return descricaoStatusPedido[s]
}
