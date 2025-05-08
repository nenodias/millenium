# coding: utf-8
from pdb import set_trace
from app import db
from sqlalchemy.sql.expression import text
from sqlalchemy.ext.declarative import declared_attr
from sqlalchemy import or_, and_, desc
from datetime import datetime
from app.utils import from_datetime_to_str

tupla_tipo_historico = (
    ('O.S.', 'Ordem de Serviço'),
    ('Orç.', 'Orçamento')
)

tupla_tipo_item = (
    ('S', 'Serviço'),
    ('F', 'Falha'),
    ('P', 'Peça')
)

tupla_origem = (
    ('', 'Selecionar'),
    (True, 'Ativo'),
    (False, 'Desativado')
)

items_colunas = ['id', 'ordem', 'tipo', 'descricao', 'quantidade', 'valor']

vistoria_colunas = ['id', 'kilometragem', 'observacao']

historico_colunas = [
    'id',
    'id_cliente',
    'id_veiculo',
    'id_tecnico',
    'numero_ordem',
    'placa',
    'sistema',
    'data',
    'tipo',
    'valor_total',
    'observacao',
    ('items', items_colunas),
    ('vistoria', vistoria_colunas)
]

cliente_colunas = [
    'id',
    'nome',
    'rg',
    'cpf',
    'endereco',
    'complemento',
    'bairro',
    'cidade',
    'cep',
    'estado',
    'pais',
    'telefone',
    'fax',
    'celular',
    'telefone_comercial',
    'fax_comercial',
    'email',
    'bip',
    'data_nascimento',
    'mes'
]

veiculo_colunas = [
    'id',
    'id_cliente',
    'id_modelo',
    'placa',
    'pais',
    'cor',
    'combustivel',
    'renavam',
    'chassi',
    'ano'
]

modelo_colunas = ['id', 'nome', 'codvei_ea', 'id_monta']

montadora_colunas = ['id', 'origem', 'nome', 'codmon_ea']

lembrete_colunas = [
    'id',
    'id_cliente',
    'id_veiculo',
    'texto',
    'data_notificacao'
]


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
            try:
                if isinstance(col, str):
                    item[col] = resolve_column(instance, col)
                else:
                    master_field = col[0]
                    detail_fields = col[1]
                    item[master_field] = []
                    item_chield = None
                    if hasattr(instance, master_field):
                        instances = getattr(instance, master_field)
                        if isinstance(instances, list):
                            for child_instance in instances:
                                item_chield = {}
                                for field in detail_fields:
                                    item_chield[field] = resolve_column(
                                        child_instance,
                                        field
                                    )
                                item[master_field].append(item_chield)
                        else:
                            item_chield = {}
                            for field in detail_fields:
                                item_chield[field] = resolve_column(
                                    instances, field
                                )
                            item[master_field] = item_chield
            except Exception as ex:
                print(ex)

        return item

    @classmethod
    def sorting_data(cls, fetch, _sort_order, _sort_direction):
        if _sort_order and _sort_direction and hasattr(cls, _sort_order):
            order = getattr(cls, _sort_order)
            if _sort_direction == 'desc':
                order = desc(order)
            fetch = fetch.order_by(order)
        elif '.' in _sort_order:
            order_list = _sort_order.split('.')
            order = None
            retorn_obj = cls
            for order_item in order_list:
                if (
                    order_item and
                    _sort_direction and
                    hasattr(retorn_obj, order_item)
                ):
                    order = getattr(retorn_obj, order_item)
                    if (
                        order and
                        hasattr(order, 'prop') and
                        hasattr(order.property, 'mapper')
                    ):
                        retorn_obj = order.prop.mapper.class_
                        fetch = fetch.join(retorn_obj)
            if order:
                if _sort_direction == 'desc':
                    order = desc(order)
                fetch = fetch.order_by(order)

        return fetch


class Cliente(MixinSerialize, db.Model):
    __tablename__ = 'cliente'

    id = db.Column(
        'id',
        db.Integer,
        primary_key=True,
        server_default=text("nextval('clientes_codigo_cliente_seq'::regclass)")
    )
    nome = db.Column('nome', db.String(60), nullable=False)
    rg = db.Column('ie_rg', db.String(16))
    cpf = db.Column('cpf', db.String(19))
    endereco = db.Column(db.String(50))
    complemento = db.Column(db.String(30))
    bairro = db.Column(db.String(30))
    cidade = db.Column(db.String(30))
    cep = db.Column(db.String(9))
    estado = db.Column(db.String(2))
    pais = db.Column(db.String(20))
    telefone = db.Column('telefone', db.String(20))
    fax = db.Column('fax', db.String(20))
    celular = db.Column(db.String(20))
    telefone_comercial = db.Column('tel_comercial', db.String(20))
    fax_comercial = db.Column('fax_comercial', db.String(20))
    email = db.Column('email', db.String(40))
    bip = db.Column('bip', db.String(30))
    data_nascimento = db.Column('data_nascimento', db.DateTime)
    mes = db.Column(db.Integer)


class Historico(MixinSerialize, db.Model):
    __tablename__ = 'historico'

    id = db.Column(
        'id',
        db.Integer,
        primary_key=True,
        server_default=text("nextval('historico_sequencia_seq'::regclass)")
    )
    id_veiculo = db.Column(
        'id_veiculo',
        db.ForeignKey('veiculo.id'),
        nullable=False
    )
    id_cliente = db.Column(
        'id_cliente',
        db.ForeignKey('cliente.id'),
        nullable=False
    )
    id_tecnico = db.Column('id_tecnico', db.ForeignKey('tecnico.id'))
    numero_ordem = db.Column('numero', db.Integer)
    placa = db.Column(db.String(8))
    sistema = db.Column(db.Integer)
    data = db.Column(db.DateTime)
    tipo = db.Column(db.String(4))
    valor_total = db.Column(db.Float(53))
    observacao = db.Column('observacao', db.String(500))

    cliente = db.relationship('Cliente')
    veiculo = db.relationship('Veiculo')
    tecnico = db.relationship('Tecnico')
    items = db.relationship('HistoricoItem', backref='historico')
    vistoria = db.relationship('Vistoria', backref='historico', uselist=False)


class HistoricoItem(MixinSerialize, db.Model):
    __tablename__ = 'historico_item'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('histitem_id_seq'::regclass)")
    )
    id_historico = db.Column(
        'id_historico',
        db.ForeignKey('historico.id'),
        nullable=False
    )
    ordem = db.Column('ordem', db.Integer, nullable=False)
    tipo = db.Column(db.String(1))
    descricao = db.Column('historico', db.String(75))
    quantidade = db.Column('quantidade', db.Integer)
    valor = db.Column(db.Float(53))


class Modelo(MixinSerialize, db.Model):
    __tablename__ = 'modelo'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('modelo_id_seq'::regclass)")
    )
    nome = db.Column('nome', db.String(40), nullable=False)
    codvei_ea = db.Column(db.Integer)
    id_monta = db.Column('id_montadora', db.ForeignKey('montadora.id'))

    montadora = db.relationship('Montadora', backref='modelos')


class Montadora(MixinSerialize, db.Model):
    __tablename__ = 'montadora'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('monta_id_seq'::regclass)")
    )
    origem = db.Column(db.String(1), nullable=False)
    nome = db.Column('nome', db.String(20), nullable=False)
    codmon_ea = db.Column(db.Integer)


class Falha(MixinSerialize, db.Model):
    __tablename__ = 'falha'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('falhas_id_seq'::regclass)")
    )
    descricao = db.Column(db.String(60))


class Peca(MixinSerialize, db.Model):
    __tablename__ = 'peca'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('cadpecas_id_seq'::regclass)")
    )
    descricao = db.Column(db.String(60), nullable=False)
    valor = db.Column(db.Float(53))


class Servico(MixinSerialize, db.Model):
    __tablename__ = 'servico'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('cadservicos_id_seq'::regclass)")
    )
    descricao = db.Column(db.String(60), nullable=False)
    valor = db.Column(db.Float(53))


class Tecnico(MixinSerialize, db.Model):
    __tablename__ = 'tecnico'

    id = db.Column(
        'id',
        db.Integer,
        primary_key=True,
        server_default=text("nextval('tecnico_codigo_tecnico_seq'::regclass)")
    )
    nome = db.Column(db.String(60), nullable=False)


class Veiculo(MixinSerialize, db.Model):
    __tablename__ = 'veiculo'

    id = db.Column(
        'id',
        db.Integer,
        primary_key=True,
        server_default=text("nextval('veiculo_codveiculo_seq'::regclass)")
    )
    id_cliente = db.Column(
        'id_cliente',
        db.ForeignKey('cliente.id')
    )
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

    def descricao(self):
        retorno = "{0} - {1}".format(self.placa, str(self.ano))
        if self.modelo is not None:
            retorno += ' - ' + self.modelo.nome
            if self.modelo.montadora is not None:
                retorno += ' - ' + self.modelo.montadora.nome
        return retorno


class Vistoria(db.Model):
    __tablename__ = 'vistoria'

    id = db.Column(
        'id_historico',
        db.ForeignKey('historico.id'),
        primary_key=True
    )
    id_carrovistoria = db.Column('id_veiculo', db.Integer)
    nivel_combustivel = db.Column('nivelcomb', db.Integer)
    kilometragem = db.Column('km', db.Float)
    tocafitas = db.Column(db.SmallInteger, server_default=text("0"))
    cd = db.Column(db.SmallInteger, server_default=text("0"))
    disqueteira = db.Column(db.SmallInteger, server_default=text("0"))
    antena = db.Column(db.SmallInteger, server_default=text("0"))
    calotas = db.Column(db.SmallInteger, server_default=text("0"))
    triangulo = db.Column(db.SmallInteger, server_default=text("0"))
    macaco = db.Column(db.SmallInteger, server_default=text("0"))
    estepe = db.Column(db.SmallInteger, server_default=text("0"))
    outro1 = db.Column(db.SmallInteger, server_default=text("0"))
    outro1_descricao = db.Column('outro1descr', db.String(20))
    outro2 = db.Column(db.SmallInteger, server_default=text("0"))
    outro2_descricao = db.Column('outro2descr', db.String(20))
    observacao = db.Column('observacao', db.String(500))


class Lembrete(MixinSerialize, db.Model):
    __tablename__ = 'lembrete'

    id = db.Column(
        db.BigInteger,
        primary_key=True,
        server_default=text("nextval('lembretes_id_seq'::regclass)")
    )
    id_veiculo = db.Column(
        'id_veiculo',
        db.ForeignKey('veiculo.id')
    )
    id_cliente = db.Column(
        'id_cliente',
        db.ForeignKey('cliente.id')
    )
    texto = db.Column(db.String(5000))
    data_notificacao = db.Column(db.DateTime)

    cliente = db.relationship('Cliente')
    veiculo = db.relationship('Veiculo')

