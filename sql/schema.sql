CREATE TABLE IF NOT EXISTS `clientes` (
  `id` ;varchar(36) NOT NULL AUTO_INCREMENT,
  `nome` varchar(100) NULL,
  `cpf` varchar(11) NULL,
  `email` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `produtos` (
  `id` varchar(36) NOT NULL AUTO_INCREMENT,
  `nome` varchar(100) NOT NULL,
  `descricao` varchar(200) NULL,
  `categoria` smallint NOT NULL,
  `preco` decimal(8,2) NOT NULL,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8; 
 
CREATE TABLE IF NOT EXISTS `pedidos` (
  `id` varchar(36) NOT NULL AUTO_INCREMENT,
  `cliente_id` varchar(36) NOT NULL,
  `valor` decimal(8,2) NOT NULL,
  `status` smallint NOT NULL,
  `data` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_pedidos_clientes_idx` (`cliente_id`),
  CONSTRAINT `fk_pedidos_clientes` FOREIGN KEY (`cliente_id`) REFERENCES `clientes` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `produtos_pedidos` (
  `pedido_id` varchar(36) NOT NULL,
  `produto_id` varchar(36) NOT NULL,
  `quantidade` int NOT NULL,
  PRIMARY KEY (`pedido_id`, `produto_id`),
  KEY `fk_pedidos_produtos_pedidos_idx` (`pedido_id`),
  KEY `fk_pedidos_produtos_produtos_idx` (`produto_id`),
  CONSTRAINT `fk_pedidos_produtos_pedidos` FOREIGN KEY (`pedido_id`) REFERENCES `pedidos` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_pedidos_produtos_produtos` FOREIGN KEY (`produto_id`) REFERENCES `produtos` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```