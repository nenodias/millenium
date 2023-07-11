CREATE TABLE falha (
    id bigserial NOT NULL primary key,
    descricao VARCHAR(60)
);

CREATE TABLE peca (
    id bigserial NOT NULL primary key,
    descricao VARCHAR(60) NOT NULL,
    valor double precision
);

CREATE TABLE servico (
    id bigserial NOT NULL primary key,
    descricao VARCHAR(60) NOT NULL,
    valor double precision
);

CREATE TABLE cliente (
    id bigserial NOT NULL primary key,
    nome VARCHAR(60) NOT NULL,
    ie_rg VARCHAR(16),
    cpf VARCHAR(19),
    endereco VARCHAR(50),
    complemento VARCHAR(30),
    bairro VARCHAR(30),
    cidade VARCHAR(30),
    cep VARCHAR(9),
    estado VARCHAR(2),
    pais VARCHAR(20),
    telefone VARCHAR(20),
    fax VARCHAR(20),
    celular VARCHAR(20),
    tel_comercial VARCHAR(20),
    fax_comercial VARCHAR(20),
    email VARCHAR(40),
    bip VARCHAR(30),
    data_nascimento timestamp without time zone,
    mes integer
);

CREATE TABLE montadora (
    id bigserial NOT NULL primary key,
    origem VARCHAR(1) NOT NULL,
    nome VARCHAR(20) NOT NULL,
    codmon_ea integer
);

CREATE TABLE modelo (
    id bigserial NOT NULL primary key,
    nome VARCHAR(40) NOT NULL,
    codvei_ea integer,
    id_montadora bigint REFERENCES montadora
);


CREATE TABLE tecnico (
    id bigserial NOT NULL primary key,
    nome VARCHAR(60) NOT NULL
);


CREATE TABLE veiculo (
    id bigserial NOT NULL primary key,
    id_cliente bigint REFERENCES cliente,
    placa VARCHAR(8) NOT NULL,
    pais VARCHAR(20),
    cor VARCHAR(20),
    combustivel VARCHAR(10),
    renavam VARCHAR(40),
    chassi VARCHAR(40),
    ano VARCHAR(4),
    id_modelo bigint
);

CREATE TABLE historico (
    id bigserial NOT NULL primary key,
    id_veiculo bigint NOT NULL REFERENCES veiculo,
    id_cliente bigint NOT NULL REFERENCES cliente,
    id_tecnico bigint REFERENCES tecnico,
    numero bigint,
    placa VARCHAR(8),
    sistema integer,
    data timestamp without time zone,
    tipo VARCHAR(4),
    valor_total double precision,
    observacao VARCHAR(500)
);

CREATE TABLE historico_item (
    id bigserial NOT NULL primary key,
    id_historico bigint NOT NULL REFERENCES historico,
    ordem bigint NOT NULL,
    tipo VARCHAR(1),
    descricao VARCHAR(75),
    quantidade integer,
    valor double precision
);

CREATE TABLE vistoria (
    id_historico bigserial NOT NULL primary key,
    id_veiculo bigint REFERENCES veiculo,
    nivelcomb integer,
    km double precision,
    tocafitas smallint DEFAULT 0,
    cd smallint DEFAULT 0,
    disqueteira smallint DEFAULT 0,
    antena smallint DEFAULT 0,
    calotas smallint DEFAULT 0,
    triangulo smallint DEFAULT 0,
    macaco smallint DEFAULT 0,
    estepe smallint DEFAULT 0,
    outro1 smallint DEFAULT 0,
    outro1descr VARCHAR(20),
    outro2 smallint DEFAULT 0,
    outro2descr VARCHAR(20),
    obs VARCHAR(500)
);

CREATE TABLE lembrete (
    id bigserial NOT NULL primary key,
    id_cliente bigint REFERENCES cliente,
    id_veiculo bigint REFERENCES veiculo,
    texto VARCHAR(5000),
    data_notificacao timestamp without time zone
);

