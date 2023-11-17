package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/felipefabricio/wonder-food/internal/entity"
)

type ClienteRepository struct {
	DB *sqlx.DB
}

func NewClienteRepository(db *sqlx.DB) *ClienteRepository {
	return &ClienteRepository{
		DB: db,
	}
}

func (r *ClienteRepository) Inserir(cliente *entity.Cliente) error {
	tx := r.DB.MustBegin()
	tx.MustExec("INSERT INTO cliente (nome, email, cpf) VALUES ($1, $2, $3)", cliente.Nome, cliente.Email, cliente.Cpf)
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
