package handler

import (
	"net/http"
)

// ListarProdutos godoc
// @Summary      Listar Produtos
// @Description  Obtém todos os Produtos cadastrados
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Success      200       {array}   entity.Produto
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /produtos [get]
func ObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CriarProduto godoc
// @Summary      Inserir Produto
// @Description  Cria um novo Produto
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        request     body      dto.ProdutoDto  true  "Request do Produto"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /produtos [post]
func InserirProduto(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
