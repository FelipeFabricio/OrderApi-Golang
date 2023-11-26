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
	"github.com/go-chi/chi/middleware"
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
	fmt.Printf("Conectado ao banco de dados na porta %s!\n", configs.DBPort)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	produtoDb := database.NewProdutoDb(db)
	produtoUseCases := usecase.NewProdutoUseCases(produtoDb)
	produtoHandler := handler.NewProdutoHandler(produtoUseCases)
	r.Route("/produtos", func(r chi.Router) {
		r.Get("/", produtoHandler.ObterTodosProdutos)
		r.Post("/", produtoHandler.InserirProduto)
		r.Get("/{categoria}", produtoHandler.ObterPorCategoria)
	})

	clienteDb := database.NewClienteDb(db)
	clienteUseCases := usecase.NewClienteUseCases(clienteDb)
	clienteHandler := handler.NewClienteHandler(clienteUseCases)
	r.Route("/clientes", func(r chi.Router) {
		r.Post("/", clienteHandler.Inserir)
		r.Get("/", clienteHandler.ObterTodos)
		r.Put("/{id}", clienteHandler.Atualizar)
	})

	pedidoDb := database.NewPedidoDb(db)
	pedidoUseCases := usecase.NewPedidoUseCases(pedidoDb)
	pedidoHandler := handler.NewPedidoHandler(pedidoUseCases)
	r.Route("/pedidos", func(r chi.Router) {
		r.Get("/", pedidoHandler.ObterPedidosEmAberto)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	fmt.Printf("Servidor iniciado na porta %s!", configs.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.WebServerPort), r)
}
