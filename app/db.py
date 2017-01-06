# -*- coding: utf-8 -*-
import time
from datetime import datetime, date
import dataset
from sqlalchemy.pool import NullPool
db = dataset.connect('sqlite:///banco.sqlite', engine_kwargs={'poolclass':NullPool} )

especialidades = db['especialidades']
medicos = db['medicos']
pacientes = db['pacientes']
agendamentos = db['agendamentos']


tupla_status = ( ('', 'Selecionar'),(True, 'Ativo'),(False, 'Desativado') )

tupla_estado = (
        ('', 'Selecionar'),
        (u'AC',u'Acre'),
        (u'AL',u'Alagoas'),
        (u'AP',u'Amapá'),
        (u'AM',u'Amazonas'),
        (u'BA',u'Bahia'),
        (u'CE',u'Ceará'),
        (u'DF',u'Distrito Federal'),
        (u'ES',u'Espírito Santo'),
        (u'GO',u'Goiás'),
        (u'MA',u'Maranhão'),
        (u'MT',u'Mato Grosso'),
        (u'MS',u'Mato Grosso do Sul'),
        (u'MG',u'Minas Gerais'),
        (u'PA',u'Pará'),
        (u'PB',u'Paraíba'),
        (u'PR',u'Paraná'),
        (u'PE',u'Pernambuco'),
        (u'PI',u'Piauí'),
        (u'RJ',u'Rio de Janeiro'),
        (u'RN',u'Rio Grande do Norte'),
        (u'RS',u'Rio Grande do Sul'),
        (u'RO',u'Rondônia'),
        (u'RR',u'Roraima'),
        (u'SC',u'Santa Catarina'),
        (u'SP',u'São Paulo'),
        (u'SE',u'Sergipe'),
        (u'TO',u'Tocantins'),
    )

tupla_plano = ( 
        ('', 'Selecionar'),
        ('SEM_COBERTURA', 'Sem Cobertura'),
        ('COBERTURA_PARCIAL', 'Cobertura parcial'),
        ('COBERTURA_TOTAL', 'Cobertura total'),
    )

tupla_area = (
    u'',
    u'Administração em Saúde',
    u'Alergia e Imunologia Pediátrica',
    u'Angiorradiologia e Cirurgia Endovascular',
    u'Atendimento ao Queimado',
    u'Cardiologia Pediátrica',
    u'Cirurgia Crânio-Maxilo-Facial',
    u'Cirurgia do trauma',
    u'Cirurgia videolaparoscópica',
    u'Citopatologia',
    u'Densitometria óssea',
    u'Dor',
    u'Ecocardiografia',
    u'ecologia',
    u'Eletrofisiologia clínica invasiva',
    u'Endocrinologia pediátrica',
    u'Endoscopia digestiva',
    u'Endoscopia ginecológica',
    u'Endoscopia respiratória',
    u'Ergometria',
    u'Foniatria',
    u'Gastroenterologia pediátrica',
    u'Hansenologia',
    u'Hematologia e hemoterapia pediátrica',
    u'Hemodinâmica e cardiologia intervencionista',
    u'Hepatologia',
    u'Infectologia hospitalar',
    u'Infectologia pediátrica',
    u'Mamografia',
    u'Medicina de urgência',
    u'Medicina do adolescente',
    u'Medicina fetal',
    u'Medicina intensiva pediátrica',
    u'Nefrologia pediátrica',
    u'Neonatologia',
    u'Neurofisiologia clínica',
    u'Neurologia pediátrica',
    u'Neurorradiologia',
    u'Nutrição parenteral e enteral',
    u'Nutrição parenteral e enteral pediátrica',
    u'Nutrologia pediátrica',
    u'Pneumologia pediátrica',
    u'Psicogeriatria',
    u'Psicoterapia',
    u'Psiquiatria da infância e adolescência',
    u'Psiquiatria forense',
    u'Radiologia intervencionista e angiorradiologia',
    u'Reumatologia pediátrica',
    u'Transplante de medula óssea',
    u'Ultrassonografia em ginecologia e obstetrícia',
)

if __name__ == '__main__':
    especialidade_values ={
        'descricao':u'',
        'status':True,

    }
    medico_values = {
        'nome':u'Médico de Teste',
        'cpf':u'111.222.444.777-35',
        'crm':u'1666/2003',
        'status':True,
        'id_especialidade':1,
        'area':u'Ecocardiografia',
        'numero':u'10-5',
        'rua':u'Rua de Teste',
        'bairro':u'Centro',
        'cidade':u'Bauru',
        'estado':u'SP',
        'telefone':u'+55 14 9 1234-5678',
    }

    paciente_values = {
        'nome':u'Paciente de Teste',
        'cpf':u'111.222.444.777-35',
        'telefone':u'+55 14 9 1234-5678',
        'plano':u'SEM_COBERTURA',
        'rua':u'Rua de Teste',
        'numero':u'10-5',
        'bairro':u'Centro',
        'cidade':u'Bauru',
        'estado':u'SP',
    }

    agendamento_values = {
        'id_medico':1,
        'id_paciente':1,
        'data':date.today(),
        'hora':datetime.strftime(datetime.now(), '%H:%M:%S'),
    }

    #results = medicos.find(_limit=10,_offset=1)
    especialidades.insert(especialidade_values)
    medicos.insert(medico_values)
    pacientes.insert(paciente_values)
    agendamentos.insert(agendamento_values)