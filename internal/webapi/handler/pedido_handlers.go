package handler

import (
	"encoding/json"
	"net/http"

	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
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
