package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felipefabricio/wonder-food/internal/dto"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
	"github.com/go-chi/chi"
)

type PedidoHandler struct {
	PedidoUseCases interfaces.PedidoUseCasesInterface
}

func NewPedidoHandler(pedidoUseCases interfaces.PedidoUseCasesInterface) *PedidoHandler {
	return &PedidoHandler{
		PedidoUseCases: pedidoUseCases,
	}
}

// ListarPedidosEmAberto godoc
// @Summary      Listar Pedido em Aberto
// @Description  Obtém todos os Pedidos que ainda não foram Finalizados
// @Tags         Pedidos
// @Accept       json
// @Produce      json
// @Success      200       {array}   dto.ObterPedidosOutputDto
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /pedidos  [get]
func (p *PedidoHandler) ObterPedidosEmAberto(w http.ResponseWriter, r *http.Request) {
	pedidos, err := p.PedidoUseCases.ObterPedidosEmAberto()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedidos)
}

// CriarPedidos godoc
// @Summary      Inserir Pedido
// @Description  Cria um novo Pedido
// @Tags         Pedidos
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CriarPedidoInputDto  true  "Dados para Cadastro do Pedido"
// @Success      201
// @Failure      400         {object}  Error
// @Failure      500         {object}  Error
// @Router       /pedidos [post]
func (p *PedidoHandler) Inserir(w http.ResponseWriter, r *http.Request) {
	var pedidoDto dto.CriarPedidoInputDto
	err := json.NewDecoder(r.Body).Decode(&pedidoDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = p.PedidoUseCases.Inserir(&pedidoDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ObterStatusPedido godoc
// @Summary      Obtém o Status do Pedido
// @Description  Consulta o Status atual do Pedido
// @Tags         Pedidos
// @Accept       json
// @Produce      json
// @Success      200            {object}  dto.ObterStatusPedidoOutputDto
// @Param        numeropedido   path      int   true  "Numero do Pedido"
// @Failure      404            {object}  Error
// @Failure      500            {object}  Error
// @Router       /pedidos/{numeropedido}  [get]
func (p *PedidoHandler) ObterStatusPedido(w http.ResponseWriter, r *http.Request) {
	numeroPedido, err := strconv.Atoi(chi.URLParam(r, "numeropedido"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pedido, err := p.PedidoUseCases.ConsultarStatusPagamento(numeroPedido)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedido)
}
