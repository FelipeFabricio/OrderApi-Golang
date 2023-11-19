package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/felipefabricio/wonder-food/configs"
	"github.com/felipefabricio/wonder-food/internal/infra/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	configs, err := configs.LoadConfig("../configs")
	if err != nil {
		panic(err)
	}

	dbConnection, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com a Base de Dados estabelecida com Sucesso!")

	queries := database.New(dbConnection)

	result, err := queries.ObterTodosProdutos(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	err = queries.InserirCliente(ctx, database.InserirClienteParams{
		ID:    uuid.New().String(),
		Nome:  sql.NullString{String: "Bebel Tuco", Valid: true},
		Email: sql.NullString{String: "felipe@uol.com", Valid: true},
		Cpf:   sql.NullString{String: "34566678900", Valid: true},
	})

	if err != nil {
		log.Fatal(err)
	}

	cliente, err := queries.ObterClientePorCpf(ctx, sql.NullString{String: "34566678900", Valid: true})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cliente)
}
