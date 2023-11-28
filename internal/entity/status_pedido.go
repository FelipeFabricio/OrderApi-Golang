package entity

type StatusPedido int

const (
	Recebido StatusPedido = iota
	EmPreparacao
	Pronto
	Finalizado
)
