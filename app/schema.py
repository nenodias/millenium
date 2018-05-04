from app import  db, ma
from app.models import Tecnico, Cliente

class ClienteSchema(ma.ModelSchema):
    class Meta:
        model = Cliente

cliente_schema = ClienteSchema()
clientes_schema = ClienteSchema(many=True)


class TecnicoSchema(ma.ModelSchema):
    class Meta:
        model = Tecnico


tecnico_schema = TecnicoSchema()
tecnicos_schema = TecnicoSchema(many=True)