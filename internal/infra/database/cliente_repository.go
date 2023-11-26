package database

import (
	"github.com/felipefabricio/wonder-food/internal/entity"
	"gorm.io/gorm"
)

type ClienteDb struct {
	DB *gorm.DB
}

func NewClienteDb(db *gorm.DB) *ClienteDb {
	return &ClienteDb{DB: db}
}

func (c *ClienteDb) ObterTodos() (*[]entity.Cliente, error) {
	var clientes *[]entity.Cliente
	err := c.DB.Find(&clientes).Error
	return clientes, err
}

func (c *ClienteDb) Inserir(cliente *entity.Cliente) error {
	return c.DB.Create(cliente).Error
}

func (c *ClienteDb) Atualizar(cliente *entity.Cliente) error {
	_, err := c.obterPorId(cliente.ID.String())
	if err != nil {
		return err
	}
	return c.DB.Save(cliente).Error
}

func (c *ClienteDb) obterPorId(id string) (*entity.Cliente, error) {
	var cliente entity.Cliente
	err := c.DB.First(&cliente, "Id = ?", id).Error
	return &cliente, err
}
