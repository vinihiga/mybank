INSERT INTO clientes (nome, limite, saldo) VALUES 
    ('a', 1000 * 100, 0),
    ('b', 0, 0),
    ('c', 10000 * 100, 0);

INSERT INTO transacoes (clienteId, tipo, valor, descricao) VALUES
    (2, 'c', 500, 'teste descricao');