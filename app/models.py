# coding: utf-8
from pdb import set_trace
from app import db
from sqlalchemy.sql.expression import text
from sqlalchemy.ext.declarative import declared_attr

class MixinSerialize():

    @classmethod
    def to_dict(cls, instance, colunas):
        item = {}
        for col in colunas:
            if hasattr(instance, col):
                item[col] = getattr(instance, col)
        return item

class Cliente(db.Model):
    __tablename__ = 'clientes'

    codigo_cliente = db.Column(db.Integer, primary_key=True, server_default=text("nextval('clientes_codigo_cliente_seq'::regclass)"))
    nome_cliente = db.Column(db.String(60), nullable=False)
    ie_rg = db.Column(db.String(16))
    cgc = db.Column(db.String(19))
    endereco = db.Column(db.String(50))
    complemento = db.Column(db.String(30))
    bairro = db.Column(db.String(30))
    cidade = db.Column(db.String(30))
    cep = db.Column(db.String(9))
    estado = db.Column(db.String(2))
    pais = db.Column(db.String(20))
    tel_res = db.Column(db.String(20))
    fax_res = db.Column(db.String(20))
    celular = db.Column(db.String(20))
    tel_com = db.Column(db.String(20))
    fax_com = db.Column(db.String(20))
    e_mail = db.Column(db.String(40))
    bip_cod = db.Column(db.String(30))
    dtnasc = db.Column(db.DateTime)
    mes = db.Column(db.Integer)

class Historico(db.Model):
    __tablename__ = 'historico'

    sequencia = db.Column(db.Integer, primary_key=True, server_default=text("nextval('historico_sequencia_seq'::regclass)"))
    codveiculo = db.Column(db.Integer, nullable=False)
    codigo_cliente = db.Column(db.ForeignKey('clientes.codigo_cliente'), nullable=False)
    tecnico = db.Column(db.ForeignKey('tecnico.codigo_tecnico'))
    nr_ordem = db.Column(db.Integer)
    placa = db.Column(db.String(8))
    sistema = db.Column(db.Integer)
    data = db.Column(db.DateTime)
    tipo = db.Column(db.String(4))
    valor_total = db.Column(db.Float(53))
    obs = db.Column(db.String(500))

    cliente = db.relationship('Cliente')
    tecnico1 = db.relationship('Tecnico')


class HistoricoItem(db.Model):
    __tablename__ = 'historico_item'

    sequencia = db.Column(db.ForeignKey('historico.sequencia'), nullable=False)
    item = db.Column(db.Integer, nullable=False)
    tipo = db.Column(db.String(1))
    historico = db.Column(db.String(75))
    qtd = db.Column(db.Integer)
    valor = db.Column(db.Float(53))
    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('histitem_id_seq'::regclass)"))

    historico1 = db.relationship('Historico')


class Modelo(db.Model):
    __tablename__ = 'modelo'

    nome_modelo = db.Column(db.String(40), nullable=False)
    codvei_ea = db.Column(db.Integer)
    id_monta = db.Column(db.ForeignKey('montadora.id'))
    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('modelo_id_seq'::regclass)"))

    montadora = db.relationship('Montadora')


class Montadora(db.Model):
    __tablename__ = 'montadora'

    origem = db.Column(db.String(1), nullable=False)
    nome_montadora = db.Column(db.String(20), nullable=False)
    codmon_ea = db.Column(db.Integer)
    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('monta_id_seq'::regclass)"))

class Falha(MixinSerialize, db.Model):
    __tablename__ = 'falhas'

    descricao = db.Column(db.String(60))
    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('falhas_id_seq'::regclass)"))

class Peca(db.Model):
    __tablename__ = 'pecas'

    descricao = db.Column(db.String(60), nullable=False)
    valor = db.Column(db.Float(53))
    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('cadpecas_id_seq'::regclass)"))


class Servico(db.Model):
    __tablename__ = 'servicos'

    descricao = db.Column(db.String(60), nullable=False)
    valor = db.Column(db.Float(53))
    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('cadservicos_id_seq'::regclass)"))


class Tecnico(db.Model):
    __tablename__ = 'tecnico'

    codigo_tecnico = db.Column(db.Integer, primary_key=True, server_default=text("nextval('tecnico_codigo_tecnico_seq'::regclass)"))
    nome = db.Column(db.String(60), nullable=False)


class Tipoitem(db.Model):
    __tablename__ = 'tipoitem'

    tipo = db.Column(db.String(1), primary_key=True)
    descricao = db.Column(db.String(15))


class Veiculo(db.Model):
    __tablename__ = 'veiculo'

    codveiculo = db.Column(db.Integer, primary_key=True, server_default=text("nextval('veiculo_codveiculo_seq'::regclass)"))
    codigo_cliente = db.Column(db.ForeignKey('clientes.codigo_cliente'))
    placa = db.Column(db.String(8), nullable=False)
    pais = db.Column(db.String(20))
    cor = db.Column(db.String(20))
    combustivel = db.Column(db.String(10))
    renavam = db.Column(db.String(40))
    chassi = db.Column(db.String(40))
    ano = db.Column(db.String(4))
    id_modelo = db.Column(db.ForeignKey('modelo.id'))

    cliente = db.relationship('Cliente')
    modelo = db.relationship('Modelo')

class Histsinal(db.Model):
    __tablename__ = 'histsinal'

    sequencia = db.Column(db.Integer, primary_key=True, nullable=False)
    sinal = db.Column(db.Integer, primary_key=True, nullable=False)
    arquivo = db.Column(db.String(25))
    descricao = db.Column(db.String(500))
    leitura = db.Column(db.String(20))

class Carrovistoria(db.Model):
    __tablename__ = 'carrovistoria'

    codigo = db.Column(db.Integer, primary_key=True, server_default=text("nextval('carrovistoria_codigo_seq'::regclass)"))
    nome = db.Column(db.String(13), nullable=False)
    arquivo = db.Column(db.String(30), nullable=False)


class Contfalha(db.Model):
    __tablename__ = 'contfalha'

    index_pk = db.Column(db.Integer, primary_key=True, nullable=False)
    continuo_fk = db.Column(db.ForeignKey('continuo.continuo_pk'), primary_key=True, nullable=False)
    codigo = db.Column(db.Integer)
    descricao = db.Column(db.String(70))
    estado = db.Column(db.String(8))

    continuo = db.relationship('Continuo')


class Contg4(db.Model):
    __tablename__ = 'contg4'

    leitura_pk = db.Column(db.Integer, primary_key=True, nullable=False)
    continuo_fk = db.Column(db.ForeignKey('continuo.continuo_pk'), primary_key=True, nullable=False)
    cocorr = db.Column(db.String(8))
    dil = db.Column(db.String(8))
    co = db.Column(db.String(8))
    co2 = db.Column(db.String(8))
    o2 = db.Column(db.String(8))
    hc = db.Column(db.String(8))
    _lambda = db.Column('lambda', db.String(8))
    rpm = db.Column(db.String(8))
    temp = db.Column(db.String(8))

    continuo = db.relationship('Continuo')


class Continuo(db.Model):
    __tablename__ = 'continuo'

    continuo_pk = db.Column(db.Integer, primary_key=True, server_default=text("nextval('continuo_continuo_pk_seq'::regclass)"))
    codigo_cliente = db.Column(db.ForeignKey('clientes.codigo_cliente'))
    codveiculo = db.Column(db.ForeignKey('veiculo.codveiculo'))
    nr_ordem = db.Column(db.Integer)
    data = db.Column(db.DateTime)

    cliente = db.relationship('Cliente')
    veiculo = db.relationship('Veiculo')


class Contparam(db.Model):
    __tablename__ = 'contparam'

    index_pk = db.Column(db.Integer, primary_key=True, nullable=False)
    continuo_fk = db.Column(db.ForeignKey('continuo.continuo_pk'), primary_key=True, nullable=False)
    tipo = db.Column(db.String(1))
    descricao = db.Column(db.String(40))
    un = db.Column(db.String(6))
    vlrlido = db.Column(db.Float(53))
    vlrlidocor = db.Column(db.String(1))
    vlrminlido = db.Column(db.Float(53))
    vlrminlidocor = db.Column(db.String(1))
    vlrmaxlido = db.Column(db.Float(53))
    vlrmaxlidocor = db.Column(db.String(1))
    vlrminimo = db.Column(db.Float(53))
    vlrmaximo = db.Column(db.Float(53))

    continuo = db.relationship('Continuo')


class Contstatu(db.Model):
    __tablename__ = 'contstatus'

    index_pk = db.Column(db.Integer, primary_key=True, nullable=False)
    continuo_fk = db.Column(db.ForeignKey('continuo.continuo_pk'), primary_key=True, nullable=False)
    descricao = db.Column(db.String(40))
    un = db.Column(db.String(6))
    vlrlido = db.Column(db.String(20))

    continuo = db.relationship('Continuo')


class Conttabg4(db.Model):
    __tablename__ = 'conttabg4'

    index_pk = db.Column(db.Integer, primary_key=True, nullable=False)
    continuo_fk = db.Column(db.ForeignKey('continuo.continuo_pk'), primary_key=True, nullable=False)
    string = db.Column(db.String(70))

    continuo = db.relationship('Continuo')


class NewPeca(db.Model):
    __tablename__ = 'new_pecas'

    ind_pecas = db.Column(db.Integer, primary_key=True, server_default=text("nextval('pecas_ind_pecas_seq'::regclass)") )
    descricao = db.Column(db.String(60), nullable=False)
    qtd = db.Column(db.Integer)
    valor = db.Column(db.Float)


class NewServico(db.Model):
    __tablename__ = 'new_servicos'

    ind_servicos = db.Column(db.Integer, primary_key=True, server_default=text("nextval('servicos_ind_servicos_seq'::regclass)"))
    descricao = db.Column(db.String(60), nullable=False)
    qtd = db.Column(db.Integer)
    valor = db.Column(db.Float)


class Notaco(db.Model):
    __tablename__ = 'notacoes'

    sequencia = db.Column(db.Integer, primary_key=True, nullable=False)
    notacao = db.Column(db.Integer, primary_key=True, nullable=False)
    top = db.Column(db.Integer)
    descricao = db.Column(db.String(30))
    esq = db.Column(db.Integer)


class Vistoria(db.Model):
    __tablename__ = 'vistoria'

    sequencia = db.Column(db.Integer, primary_key=True, server_default=text("nextval('vistoria_sequencia_seq'::regclass)"))
    carrovistoria = db.Column(db.Integer)
    nivelcomb = db.Column(db.Integer)
    kilometragem = db.Column(db.Float)
    tocafitas = db.Column(db.SmallInteger, server_default=text("0"))
    cd = db.Column(db.SmallInteger, server_default=text("0"))
    disqueteira = db.Column(db.SmallInteger, server_default=text("0"))
    antena = db.Column(db.SmallInteger, server_default=text("0"))
    calotas = db.Column(db.SmallInteger, server_default=text("0"))
    triangulo = db.Column(db.SmallInteger, server_default=text("0"))
    macaco = db.Column(db.SmallInteger, server_default=text("0"))
    estepe = db.Column(db.SmallInteger, server_default=text("0"))
    outro1 = db.Column(db.SmallInteger, server_default=text("0"))
    outro1descr = db.Column(db.String(20))
    outro2 = db.Column(db.SmallInteger, server_default=text("0"))
    outro2descr = db.Column(db.String(20))
    obs = db.Column(db.String(500))
