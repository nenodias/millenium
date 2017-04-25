# coding: utf-8
from pdb import set_trace
from app import db
from sqlalchemy.sql.expression import text
from sqlalchemy.ext.declarative import declared_attr
from sqlalchemy import or_, and_, desc
from datetime import datetime
from app.utils import from_datetime_to_str

tupla_tipo_historico = ( ('O.S.', 'Ordem de Serviço'), ('Orç.', 'Orçamento') )

tupla_tipo_item = ( ('S', 'Serviço'),('F', 'Falha'), ('P', 'Peça') )

tupla_origem = ( ('', 'Selecionar'),(True, 'Ativo'),(False, 'Desativado') )

items_colunas = ['id','ordem','tipo','descricao','quantidade','valor']

vistoria_colunas = ['id','kilometragem','observacao']

historico_colunas = [ 'id', 'id_cliente', 'id_veiculo', 'id_tecnico', 'numero_ordem', 'placa', 'sistema', 'data', 'tipo', 'valor_total', 'observacao', ('items', items_colunas ), ('vistoria', vistoria_colunas)]

cliente_colunas = [ 'id','nome','rg','cpf','endereco','complemento','bairro','cidade','cep','estado','pais','telefone','fax','celular','telefone_comercial','fax_comercial','email','bip','data_nascimento','mes' ]

veiculo_colunas = [ 'id','id_cliente','id_modelo','placa','pais','cor','combustivel','renavam','chassi','ano']

modelo_colunas = [ 'id', 'nome', 'codvei_ea', 'id_monta' ]

montadora_colunas = [ 'id', 'origem', 'nome', 'codmon_ea' ]

lembrete_colunas = ['id', 'id_cliente', 'id_veiculo', 'texto', 'data_notificacao']

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
                        instances =  getattr(instance, master_field)
                        if isinstance(instances,list):
                            for child_instance in instances:
                                item_chield = {}
                                for field in detail_fields:
                                    item_chield[field] = resolve_column(child_instance, field)
                                item[master_field].append( item_chield )
                        else:
                            item_chield = {}
                            for field in detail_fields:
                                item_chield[field] = resolve_column(instances, field)
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
                if order_item and _sort_direction and hasattr(retorn_obj, order_item):
                    order = getattr(retorn_obj, order_item)
                    if order and hasattr(order,'prop') and hasattr(order.property,'mapper'):
                        retorn_obj = order.prop.mapper.class_
                        fetch = fetch.join(retorn_obj)
            if order:
                if _sort_direction == 'desc':
                    order = desc(order)
                fetch = fetch.order_by(order)


        return fetch

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
    items = db.relationship('HistoricoItem', backref='historico')
    vistoria = db.relationship('Vistoria', backref='historico',uselist=False)


class HistoricoItem(MixinSerialize, db.Model):
    __tablename__ = 'historico_item'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('histitem_id_seq'::regclass)"))
    id_historico = db.Column('sequencia', db.ForeignKey('historico.sequencia'), nullable=False)
    ordem = db.Column('item',db.Integer, nullable=False)
    tipo = db.Column(db.String(1))
    descricao = db.Column('historico',db.String(75))
    quantidade = db.Column('qtd',db.Integer)
    valor = db.Column(db.Float(53))

    #historico = db.relationship('Historico', backref='items', lazy="dynamic")


class Modelo(MixinSerialize, db.Model):
    __tablename__ = 'modelo'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('modelo_id_seq'::regclass)"))
    nome = db.Column('nome_modelo', db.String(40), nullable=False)
    codvei_ea = db.Column(db.Integer)
    id_monta = db.Column(db.ForeignKey('montadora.id'))

    montadora = db.relationship('Montadora', backref='modelos')


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

class Vistoria(db.Model):
    __tablename__ = 'vistoria'

    #id = db.Column('sequencia', db.Integer, primary_key=True, server_default=text("nextval('vistoria_sequencia_seq'::regclass)"))
    id = db.Column('sequencia', db.ForeignKey('historico.sequencia'), primary_key=True)
    id_carrovistoria = db.Column('carrovistoria',db.Integer)
    nivel_combustivel = db.Column('nivelcomb',db.Integer)
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
    outro1_descricao = db.Column('outro1descr',db.String(20))
    outro2 = db.Column(db.SmallInteger, server_default=text("0"))
    outro2_descricao = db.Column('outro2descr',db.String(20))
    observacao = db.Column('obs',db.String(500))

class Lembrete(MixinSerialize, db.Model):
    __tablename__ = 'lembretes'

    id = db.Column(db.BigInteger, primary_key=True, server_default=text("nextval('lembretes_id_seq'::regclass)"))
    id_veiculo = db.Column('id_veiculo', db.ForeignKey('veiculo.codveiculo'))
    id_cliente = db.Column('id_cliente', db.ForeignKey('clientes.codigo_cliente'))
    texto = db.Column(db.String(5000))
    data_notificacao = db.Column(db.DateTime)

    cliente = db.relationship('Cliente')
    veiculo = db.relationship('Veiculo')


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
