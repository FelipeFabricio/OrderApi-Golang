package entity

//go:generate stringer -type=CategoriaProduto
type CategoriaProduto int

// TODO: Criar validação para saber se o valor informado está no range do enum
const (
	Lanche CategoriaProduto = 1 + iota
	Bebida
	Sobremesa
	Acompanhamento
)
