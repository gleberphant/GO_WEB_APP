-- LIMPAR DATABASE
DROP TABLE IF EXISTS produtos;

-- CRIAR AS TABELAS
CREATE TABLE IF NOT EXISTS produtos (
    id INT NOT NULL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    preco REAL,
    quantidade INTEGER
);

-- Insert sample data into produtos table --/
INSERT INTO produtos (id, nome, descricao, preco, quantidade) VALUES
(1, 'Notebook', 'Notebook Dell Inspiron 15', 3500.00, 10),
(2, 'Smartphone', 'Smartphone Samsung Galaxy S21', 2500.00, 20),
(3, 'Tablet', 'Tablet Apple iPad Air', 3000.00, 15);