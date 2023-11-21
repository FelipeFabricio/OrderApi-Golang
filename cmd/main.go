package main

import (
	"fmt"
	"net/http"

	"github.com/felipefabricio/wonder-food/configs"
	_ "github.com/felipefabricio/wonder-food/docs"
	"github.com/felipefabricio/wonder-food/internal/infra/database"
	"github.com/felipefabricio/wonder-food/internal/usecase"
	"github.com/felipefabricio/wonder-food/internal/webapi/handler"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	configs, err := configs.LoadConfig("../configs")
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado ao banco de dados na porta 3306!")

	produtoRepository := database.NewProdutoDb(db)
	produtoUseCases := usecase.NewProdutoUseCases(produtoRepository)
	produtoHandler := handler.NewProdutoHandler(produtoUseCases)

	r := chi.NewRouter()
	r.Route("/produtos", func(r chi.Router) {
		r.Get("/", produtoHandler.ObterTodosProdutos)
		r.Post("/", produtoHandler.InserirProduto)
		r.Get("/{categoria:int}", produtoHandler.ObterPorCategoria)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	fmt.Println("Servidor iniciado na porta 8000!")
	http.ListenAndServe(":8000", r)
}
