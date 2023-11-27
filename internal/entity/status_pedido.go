package entity

//go:generate stringer -type=StatusPedido
type StatusPedido int

const (
	Recebido StatusPedido = iota
	EmPreparacao
	Pronto
	Finalizado
)
