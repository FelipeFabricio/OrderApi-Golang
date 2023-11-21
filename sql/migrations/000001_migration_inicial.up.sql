CREATE TABLE IF NOT EXISTS `clientes` (
  `id` varchar(36) NOT NULL,
  `nome` varchar(100) NULL,
  `cpf` varchar(11) NULL,
  `email` varchar(256) NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `clientes` (`id`,`nome`, `cpf`, `email`) VALUES
('250cc130-f258-4e32-abdd-5a457888c513', 'João da Silva', '93880743002', 'joao@gmail.com');

CREATE TABLE IF NOT EXISTS `produtos` (
  `id` varchar(36) NOT NULL,
  `nome` varchar(100) NOT NULL,
  `descricao` varchar(200) NULL,
  `categoria` smallint NOT NULL,
  `valor` decimal(8,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `produtos` (`id`, `nome`, `descricao`, `categoria`, `valor`) VALUES
('99d8b405-378f-449a-8636-c431fa38debb', 'X-Bacon', 'Pão, Hambúrguer Bovino e Bacon duplo', 1, 29.90),
('d12dc279-a731-4ae3-8713-5f655bfe966e', 'Suco de Laranja', 'Suco da fruta!!', 2, 10.90),
('9e40d3cc-7325-46c1-a5a2-8d770bdd22a0', 'Pudim', 'O lendário Pudim de leite', 3, 9.99),
('be484c18-8f83-4a76-a5c4-f1d2f98cce94', 'Batata Frita Média', 'Batata frita clássica', 4, 8.50);

CREATE TABLE IF NOT EXISTS `pedidos` (
  `id` varchar(36) NOT NULL,
  `cliente_id` varchar(36) NOT NULL,
  `valor` decimal(8,2) NOT NULL,
  `status` smallint NOT NULL,
  `data` datetime  DEFAULT now(),
  `numeroPedido` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_pedidos_clientes_idx` (`cliente_id`),
  CONSTRAINT `fk_pedidos_clientes` FOREIGN KEY (`cliente_id`) REFERENCES `clientes` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `pedidos` (`id`, `cliente_id`, `valor`, `status`, `numeroPedido`) VALUES
('1a367798-fee1-49cd-a686-fc3e04577e2e', '250cc130-f258-4e32-abdd-5a457888c513', 59.29, 3, 1);

CREATE TABLE IF NOT EXISTS `produtos_pedidos` (
  `pedido_id` varchar(36) NOT NULL,
  `produto_id` varchar(36) NOT NULL,
  `quantidade` int NOT NULL,
  CONSTRAINT `fk_pedidos_produtos_pedidos` FOREIGN KEY (`pedido_id`) REFERENCES `pedidos` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_pedidos_produtos_produtos` FOREIGN KEY (`produto_id`) REFERENCES `produtos` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `produtos_pedidos` (`pedido_id`, `produto_id`, `quantidade`) VALUES
('1a367798-fee1-49cd-a686-fc3e04577e2e', '99d8b405-378f-449a-8636-c431fa38debb', 1),
('1a367798-fee1-49cd-a686-fc3e04577e2e', 'd12dc279-a731-4ae3-8713-5f655bfe966e', 1),
('1a367798-fee1-49cd-a686-fc3e04577e2e', '9e40d3cc-7325-46c1-a5a2-8d770bdd22a0', 1),
('1a367798-fee1-49cd-a686-fc3e04577e2e', 'be484c18-8f83-4a76-a5c4-f1d2f98cce94', 1);
