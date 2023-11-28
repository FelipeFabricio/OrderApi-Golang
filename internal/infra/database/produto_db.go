package database

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	"gorm.io/gorm"
)

type ProdutoDb struct {
	DB *gorm.DB
}

func NewProdutoDb(db *gorm.DB) *ProdutoDb {
	return &ProdutoDb{DB: db}
}

func (p *ProdutoDb) Inserir(produto *entity.Produto) error {
	return p.DB.Create(produto).Error
}

func (p *ProdutoDb) ObterTodos() (*[]entity.Produto, error) {
	var produtos []entity.Produto
	err := p.DB.Find(&produtos).Error
	return &produtos, err
}

func (p *ProdutoDb) ObterPorCategoria(categoria entity.CategoriaProduto) (*[]entity.Produto, error) {
	var produto []entity.Produto
	err := p.DB.Find(&produto, "Categoria = ?", categoria).Error
	return &produto, err
}

func (p *ProdutoDb) Atualizar(produto *entity.Produto) error {
	_, err := p.ObterPorId(produto.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(produto).Error
}

func (p *ProdutoDb) Deletar(id string) error {
	produto, err := p.ObterPorId(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(produto).Error
}

func (p *ProdutoDb) ObterPorId(id string) (*entity.Produto, error) {
	var produto entity.Produto
	err := p.DB.First(&produto, "Id = ?", id).Error
	return &produto, err
}
