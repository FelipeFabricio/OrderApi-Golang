package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/felipefabricio/wonder-food/configs"
	"github.com/felipefabricio/wonder-food/internal/infra/database"
	"github.com/felipefabricio/wonder-food/internal/webapi/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// @title           WonderFood API
// @version         1.0
// @description     API para o sistema de pedidos da WonderFood
// @termsOfService  http://swagger.io/terms/

// @contact.name   Felipe Fabricio
// @contact.url    https://www.linkedin.com/in/felipefabricio/
// @contact.email  ff.oliveira32@gmail.com

// @host      localhost:8000
// @BasePath  /
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

	router.HandleRequests()

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
