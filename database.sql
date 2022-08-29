ALTER TABLE clientes RENAME COLUMN codigo_cliente TO id;
ALTER TABLE clientes RENAME COLUMN nome_cliente TO nome;
ALTER TABLE clientes RENAME COLUMN cgc TO cpf;
ALTER TABLE clientes RENAME COLUMN tel_res TO telefone;
ALTER TABLE clientes RENAME COLUMN fax_res TO fax;
ALTER TABLE clientes RENAME COLUMN tel_com TO tel_comercial;
ALTER TABLE clientes RENAME COLUMN fax_com TO fax_comercial;
ALTER TABLE clientes RENAME COLUMN e_mail TO email;
ALTER TABLE clientes RENAME COLUMN bip_cod TO bip;
ALTER TABLE clientes RENAME COLUMN dtnasc TO data_nascimento;
ALTER TABLE clientes ALTER COLUMN id to bigint

ALTER TABLE historico RENAME COLUMN sequencia TO id;
ALTER TABLE historico RENAME COLUMN codveiculo TO id_veiculo;
ALTER TABLE historico RENAME COLUMN codigo_cliente TO id_cliente;
ALTER TABLE historico RENAME COLUMN tecnico TO id_tecnico;
ALTER TABLE historico RENAME COLUMN nr_ordem TO numero;
ALTER TABLE historico RENAME COLUMN obs TO observacao;
ALTER TABLE historico ALTER COLUMN id to bigint
ALTER TABLE historico ALTER COLUMN id_veiculo to bigint
ALTER TABLE historico ALTER COLUMN id_cliente to bigint
ALTER TABLE historico ALTER COLUMN id_tecnico to bigint
ALTER TABLE historico ALTER COLUMN numero to bigint

ALTER TABLE historico_item RENAME COLUMN sequencia TO id_historico;
ALTER TABLE historico_item RENAME COLUMN qtd TO quantidade;
ALTER TABLE historico_item RENAME COLUMN item TO ordem;
ALTER TABLE historico_item ALTER COLUMN id_historico to bigint

ALTER TABLE modelo RENAME COLUMN id_monta TO id_montadora;
ALTER TABLE modelo RENAME COLUMN nome_modelo TO nome;

ALTER TABLE montadora RENAME COLUMN nome_montadora TO nome;

ALTER TABLE tecnico RENAME COLUMN codigo_tecnico TO id;
ALTER TABLE tecnico ALTER COLUMN id to bigint

ALTER TABLE veiculo RENAME COLUMN codveiculo TO id;
ALTER TABLE veiculo RENAME COLUMN codigo_cliente TO id_cliente;
ALTER TABLE veiculo ALTER COLUMN id to bigint
ALTER TABLE veiculo ALTER COLUMN id_cliente to bigint

ALTER TABLE vistoria RENAME COLUMN sequencia TO id_historico;
ALTER TABLE vistoria RENAME COLUMN carrovistoria TO id_veiculo;
ALTER TABLE vistoria RENAME COLUMN kilometragem TO km;
ALTER TABLE vistoria RENAME COLUMN obs TO observacao;
ALTER TABLE vistoria ALTER COLUMN id_historico to bigint
ALTER TABLE vistoria ALTER COLUMN id_veiculo to bigint

-- RENAMING TABLES FOR STANDARDIZE
ALTER TABLE clientes RENAME TO cliente
ALTER TABLE falhas RENAME TO falha
ALTER TABLE lembretes RENAME TO lembrete
ALTER TABLE servicos RENAME TO servicos
ALTER TABLE pecas RENAME TO peca