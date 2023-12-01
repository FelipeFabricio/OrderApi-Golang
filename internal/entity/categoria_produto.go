package entity

type CategoriaProduto int

const (
	Lanche CategoriaProduto = 1 + iota
	Bebida
	Sobremesa
	Acompanhamento
)

var descricaoCategoriaProduto = [...]string{
	"",
	"Lanche",
	"Bebida",
	"Sobremesa",
	"Acompanhamento",
}

func (c CategoriaProduto) ObterDescricaoCategoriaProduto() string {
	if c <= 0 || int(c) >= len(descricaoCategoriaProduto) {
		return "Desconhecido"
	}
	return descricaoCategoriaProduto[c]
}
