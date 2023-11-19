-- name: InserirProduto :exec
INSERT INTO Produtos (Id, Nome, Valor, Descricao, Categoria) 
VALUES (?, ?, ?, ?, ?);

-- name: ObterTodosProdutos :many
SELECT * FROM Produtos;

-- name: ObterProdutosPorCategoria :many
SELECT * FROM Produtos WHERE Categoria = ?;

-- name: AtualizarProduto :execresult
UPDATE Produtos SET Nome = ?, Valor = ?, Descricao = ?, Categoria = ? 
WHERE Id = ?;

-- name: DeletarProduto :exec
DELETE FROM Produtos WHERE Id = ?;

-- name: InserirCliente :exec
INSERT INTO Clientes (Id, Nome, Email, Cpf)
VALUES (?, ?, ?, ?);

-- name: AtualizarCliente :execresult
UPDATE Clientes SET Nome = ?, Email = ?, Cpf = ?
WHERE Id = ?;                              

-- name: ObterClientePorCpf :one
SELECT * FROM Clientes WHERE Cpf = ?;

-- name: ObterPedidoPorNumeroPedido :one
SELECT * FROM Pedidos WHERE NumeroPedido = ?;

-- name: DeletarPedido :exec
DELETE FROM Pedidos WHERE NumeroPedido = ?;













































