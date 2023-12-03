package entity

type StatusPedido int

const (
	AguardandoPagamento StatusPedido = iota
	Recebido
	EmPreparacao
	Pronto
	Finalizado
)

var descricaoStatusPedido = [...]string{
	"Aguardando pagamento",
	"Recebido",
	"Em preparação",
	"Pronto",
	"Finalizado",
}

func (s StatusPedido) ObterDescricaoStatusPedido() string {
	if s < 0 || int(s) >= len(descricaoStatusPedido) {
		return "Desconhecido"
	}
	return descricaoStatusPedido[s]
}
