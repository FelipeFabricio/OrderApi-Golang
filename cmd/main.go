package main

import (
	"fmt"
	"log"

	"github.com/felipefabricio/wonder-food/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	configs, err := configs.LoadConfig("../configs")
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com a Base de Dados estabelecida com Sucesso!")
}
