package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
	"github.com/go-chi/chi"
)

type ProdutoHandler struct {
	produtoUseCases interfaces.ProdutoUseCasesInterface
}

func NewProdutoHandler(useCases interfaces.ProdutoUseCasesInterface) *ProdutoHandler {
	return &ProdutoHandler{produtoUseCases: useCases}
}

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
func (p *ProdutoHandler) ObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	produtos, err := p.produtoUseCases.ObterTodosProdutos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produtos)
}

// ListarProdutosPorCategoria godoc
// @Summary      Listar Produto por Categoria de Produto
// @Description  Obtém todos os Produtos cadastrados por Categoria de Produto
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        categoria path      int  true  "Categoria Produto"
// @Success      200       {array}   entity.Produto
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /produtos/{categoria} [get]
func (p *ProdutoHandler) ObterPorCategoria(w http.ResponseWriter, r *http.Request) {
	//TODO: Validar se o parâmetro é um número e converter para um Enum
	_, err := strconv.Atoi(chi.URLParam(r, "categoria"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//produtos, err := p.produtoUseCases.ObterPorCategoria(categoriaProduto)
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
func (p *ProdutoHandler) InserirProduto(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
