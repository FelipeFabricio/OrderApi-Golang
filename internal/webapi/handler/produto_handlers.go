package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felipefabricio/wonder-food/internal/entity"
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
		w.WriteHeader(http.StatusNotFound)
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
// @Failure      400       {object}  Error
// @Failure      500       {object}  Error
// @Router       /produtos/{categoria} [get]
func (p *ProdutoHandler) ObterPorCategoria(w http.ResponseWriter, r *http.Request) {
	categoria, err := strconv.Atoi(chi.URLParam(r, "categoria"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	produtos, err := p.produtoUseCases.ObterPorCategoria(entity.CategoriaProduto(categoria))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//TODO: Deixar setado para todos os endpoints respoderem como json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produtos)
}

// CriarProduto godoc
// @Summary      Inserir Produto
// @Description  Cria um novo Produto
// @Tags         Produtos
// @Accept       json
// @Produce      json
// @Param        request     body      dto.ProdutoDto  true  "Request do Produto"
// @Success      201
// @Failure      400       {object}  Error
// @Failure      500         {object}  Error
// @Router       /produtos [post]
func (p *ProdutoHandler) InserirProduto(w http.ResponseWriter, r *http.Request) {
	var produto entity.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = p.produtoUseCases.Inserir(&produto)
	if err != nil {
		//TODO: Criar objeto de erro para ser mostrado no swagger
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
