package router

import (
	"net/http"

	_ "github.com/felipefabricio/wonder-food/docs"
	"github.com/felipefabricio/wonder-food/internal/webapi/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func HandleRequests() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/produtos", func(r chi.Router) {
		r.Get("/", handler.ObterTodosProdutos)
		r.Post("/", handler.InserirProduto)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
