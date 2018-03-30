from app import  db, ma
from app.models import Tecnico

class TecnicoSchema(ma.ModelSchema):
    class Meta:
        model = Tecnico


tecnico_schema = TecnicoSchema()
tecnicos_schema = TecnicoSchema(many=True)