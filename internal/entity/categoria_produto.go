package entity

type CategoriaProduto int

const (
	Lanche CategoriaProduto = 1 + iota
	Bebida
	Sobremesa
	Acompanhamento
)
