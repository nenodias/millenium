from app import  db, ma
from app.models import Historico, HistoricoItem, Falha, Tecnico, Cliente, Peca, Servico, Modelo, Montadora, Veiculo, Lembrete

class FalhaSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Falha
        load_instance = True

falha_schema = FalhaSchema()
falhas_schema = FalhaSchema(many=True)

###

class ClienteSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Cliente
        load_instance = True

cliente_schema = ClienteSchema()
clientes_schema = ClienteSchema(many=True)

###

class TecnicoSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Tecnico
        load_instance = True


tecnico_schema = TecnicoSchema()
tecnicos_schema = TecnicoSchema(many=True)

###

class PecaSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Peca
        load_instance = True


peca_schema = PecaSchema()
pecas_schema = PecaSchema(many=True)

###

class ServicoSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Servico
        load_instance = True


servico_schema = ServicoSchema()
servicos_schema = ServicoSchema(many=True)

###

class ModeloSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Modelo
        load_instance = True


modelo_schema = ModeloSchema()
modelos_schema = ModeloSchema(many=True)

###

class MontadoraSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Montadora
        load_instance = True


montadora_schema = MontadoraSchema()
montadoras_schema = MontadoraSchema(many=True)

###

class VeiculoSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Veiculo
        load_instance = True


veiculo_schema = VeiculoSchema()
veiculos_schema = VeiculoSchema(many=True)

###

class LembreteSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Lembrete
        load_instance = True


lembrete_schema = LembreteSchema()
lembretes_schema = LembreteSchema(many=True)

class HistoricoItemSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = HistoricoItem
        load_instance = True

historico_item_schema = HistoricoItemSchema()
historico_items_schema = HistoricoItemSchema(many=True)


class HistoricoSchema(ma.SQLAlchemyAutoSchema):
    items = ma.Nested(historico_items_schema)
    cliente = ma.Nested(cliente_schema)
    veiculo = ma.Nested(veiculo_schema)
    tecnico = ma.Nested(tecnico_schema)
    class Meta:
        model = Historico
        load_instance = True

historico_schema = HistoricoSchema()
historicos_schema = HistoricoSchema(many=True)

###