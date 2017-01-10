# coding: utf-8
from pdb import set_trace
from app import db
from sqlalchemy.sql.expression import text
from sqlalchemy.ext.declarative import declared_attr
from sqlalchemy import or_, and_
from datetime import datetime
from app.utils import from_datetime_to_str

class MixinSerialize():

    @classmethod
    def to_dict(cls, instance, colunas):
        item = {}
        def resolve_column(inst, colu):
            if hasattr(inst, colu):
                valor = getattr(inst, colu)
                if isinstance(valor, datetime):
                    valor = from_datetime_to_str(valor)
                if valor:
                    return valor
                else:
                    return ''
        for col in colunas:
            if isinstance(col, str):
                item[col] = resolve_column(instance, col)
            else:
                childs = []
                master_field = col[0]
                detail_fields = col[1]
                item[master_field] = childs
                if hasattr(instance, master_field):
                    instances = getattr(instance, master_field)
                    for child_instance in instances:
                        item_chield = {}
                        for field in detail_fields:
                            item_chield[field] = resolve_column(child_instance, field)
                        childs.append( item_chield )


        return item

class Cliente(MixinSerialize, db.Model):
    __tablename__ = 'clientes'

    id = db.Column('codigo_cliente',db.Integer, primary_key=True, server_default=text("nextval('clientes_codigo_cliente_seq'::regclass)"))
    nome = db.Column('nome_cliente',db.String(60), nullable=False)
    rg = db.Column('ie_rg',db.String(16))
    cpf = db.Column('cgc',db.String(19))
    endereco = db.Column(db.String(50))
    complemento = db.Column(db.String(30))
    bairro = db.Column(db.String(30))
    cidade = db.Column(db.String(30))
    cep = db.Column(db.String(9))
    estado = db.Column(db.String(2))
    pais = db.Column(db.String(20))
    telefone = db.Column('tel_res',db.String(20))
    fax = db.Column('fax_res',db.String(20))
    celular = db.Column(db.String(20))
    telefone_comercial = db.Column('tel_com',db.String(20))
    fax_comercial = db.Column('fax_com',db.String(20))
    email = db.Column('e_mail',db.String(40))
    bip = db.Column('bip_cod',db.String(30))
    data_nascimento = db.Column('dtnasc',db.DateTime)
    mes = db.Column(db.Integer)

class Historico(MixinSerialize, db.Model):
    __tablename__ = 'historico'

    id = db.Column('sequencia', db.Integer, primary_key=True, server_default=text("nextval('historico_sequencia_seq'::regclass)"))
    id_veiculo = db.Column('codveiculo',db.ForeignKey('veiculo.codveiculo'), nullable=False)
    id_cliente = db.Column('codigo_cliente',db.ForeignKey('clientes.codigo_cliente'), nullable=False)
    id_tecnico = db.Column('tecnico',db.ForeignKey('tecnico.codigo_tecnico'))
    numero_ordem = db.Column('nr_ordem',db.Integer)
    placa = db.Column(db.String(8))
    sistema = db.Column(db.Integer)
    data = db.Column(db.DateTime)
    tipo = db.Column(db.String(4))
    valor_total = db.Column(db.Float(53))
    observacao = db.Column('obs',db.String(500))

    cliente = db.relationship('Cliente')
    veiculo = db.relationship('Veiculo')
    tecnico = db.relationship('Tecnico')


class HistoricoItem(MixinSerialize, db.Model):
    __tablename__ = 'historico_item'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('histitem_id_seq'::regclass)"))
    id_historico = db.Column('sequencia', db.ForeignKey('historico.sequencia'), nullable=False)
    ordem = db.Column('item',db.Integer, nullable=False)
    tipo = db.Column(db.String(1))
    descricao = db.Column('historico',db.String(75))
    quantidade = db.Column('qtd',db.Integer)
    valor = db.Column(db.Float(53))    

    historico = db.relationship('Historico', backref='items')


class Modelo(MixinSerialize, db.Model):
    __tablename__ = 'modelo'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('modelo_id_seq'::regclass)"))
    nome = db.Column('nome_modelo', db.String(40), nullable=False)
    codvei_ea = db.Column(db.Integer)
    id_monta = db.Column(db.ForeignKey('montadora.id'))

    montadora = db.relationship('Montadora')


class Montadora(MixinSerialize, db.Model):
    __tablename__ = 'montadora'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('monta_id_seq'::regclass)"))
    origem = db.Column(db.String(1), nullable=False)
    nome = db.Column('nome_montadora',db.String(20), nullable=False)
    codmon_ea = db.Column(db.Integer)

class Falha(MixinSerialize, db.Model):
    __tablename__ = 'falhas'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('falhas_id_seq'::regclass)"))
    descricao = db.Column(db.String(60))

class Peca(MixinSerialize, db.Model):
    __tablename__ = 'pecas'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('cadpecas_id_seq'::regclass)"))
    descricao = db.Column(db.String(60), nullable=False)
    valor = db.Column(db.Float(53))


class Servico(MixinSerialize, db.Model):
    __tablename__ = 'servicos'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('cadservicos_id_seq'::regclass)"))
    descricao = db.Column(db.String(60), nullable=False)
    valor = db.Column(db.Float(53))


class Tecnico(MixinSerialize, db.Model):
    __tablename__ = 'tecnico'

    id = db.Column('codigo_tecnico',db.Integer, primary_key=True, server_default=text("nextval('tecnico_codigo_tecnico_seq'::regclass)"))
    nome = db.Column(db.String(60), nullable=False)


class Veiculo(MixinSerialize, db.Model):
    __tablename__ = 'veiculo'

    id = db.Column('codveiculo',db.Integer, primary_key=True, server_default=text("nextval('veiculo_codveiculo_seq'::regclass)"))
    id_cliente = db.Column('codigo_cliente',db.ForeignKey('clientes.codigo_cliente'))
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


class Tipoitem(MixinSerialize, db.Model):
    __tablename__ = 'tipoitem'

    tipo = db.Column(db.String(1), primary_key=True)
    descricao = db.Column(db.String(15))

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
