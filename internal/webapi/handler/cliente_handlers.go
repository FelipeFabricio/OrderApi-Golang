package handler

import (
	"encoding/json"
	"net/http"

	"github.com/felipefabricio/wonder-food/internal/entity"
	interfaces "github.com/felipefabricio/wonder-food/internal/entity/interfaces"
	"github.com/go-chi/chi"
)

type ClienteHandler struct {
	clienteUseCases interfaces.ClienteUseCasesInterface
}

func NewClienteHandler(useCases interfaces.ClienteUseCasesInterface) *ClienteHandler {
	return &ClienteHandler{clienteUseCases: useCases}
}

// CriarCliente godoc
// @Summary      Inserir Cliente
// @Description  Cadastra um novo Cliente
// @Tags         Clientes
// @Accept       json
// @Produce      json
// @Param        request     body      dto.ClienteDto  true  "Dados para cadastro do Cliente"
// @Success      201
// @Failure      400         {object}  Error
// @Failure      500         {object}  Error
// @Router       /clientes   [post]
func (c *ClienteHandler) Inserir(w http.ResponseWriter, r *http.Request) {
	var cliente entity.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.clienteUseCases.Inserir(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// ListarTodosClientes godoc
// @Summary      Listar todos os Clientes
// @Description  Obtém todos os Clientes cadastrados no sistema
// @Tags         Clientes
// @Accept       json
// @Produce      json
// @Success      200        {array}   entity.Cliente
// @Failure      404        {object}  Error
// @Failure      500        {object}  Error
// @Router       /clientes  [get]
func (c *ClienteHandler) ObterTodos(w http.ResponseWriter, r *http.Request) {
	clientes, err := c.clienteUseCases.ObterTodos()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clientes)
}

// AtualizarCliente godoc
// @Summary      Atualizar Cliente
// @Description  Atualiza os Dados de um Cliente
// @Tags         Clientes
// @Accept       json
// @Produce      json
// @Param        id              path      string          true  "Id do Cliente"
// @Param        request         body      dto.ClienteDto  true  "Dados para atualizar o Cliente"
// @Success      200             {object}  entity.Cliente
// @Failure      404             {object}  Error
// @Failure      500             {object}  Error
// @Router       /clientes/{id}  [put]
func (c *ClienteHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	var cliente entity.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	if id != cliente.ID.String() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.clienteUseCases.Atualizar(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
