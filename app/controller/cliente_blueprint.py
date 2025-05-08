# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (
    Blueprint,
    render_template,
    request,
    redirect,
    url_for,
    flash,
    jsonify,
    render_template,
    Response
)
from app.authentication import auth_require
from app import  db
from app.utils import to_int_or_none, from_str_to_datetime_or_none
from app.models import Cliente, or_, and_, desc, cliente_colunas

cliente_blueprint = Blueprint('cliente', __name__)


@cliente_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _nome = request.args.get('nome', '')
    _telefone = request.args.get('telefone', '')
    _celular = request.args.get('celular', '')
    contexto['model'] = {
        'nome':_nome,
        'telefone': _telefone,
        'celular':_celular
    }
    return render_template('cliente/consulta.html', **contexto)

@cliente_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@cliente_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        nome = request.form.get('nome')
        rg = request.form.get('rg')
        cpf = request.form.get('cpf')
        endereco = request.form.get('endereco')
        complemento = request.form.get('complemento')
        bairro = request.form.get('bairro')
        cidade = request.form.get('cidade')
        cep = request.form.get('cep')
        estado = request.form.get('estado')
        pais = request.form.get('pais')
        telefone = request.form.get('telefone')
        fax = request.form.get('fax')
        celular = request.form.get('celular')
        telefone_comercial = request.form.get('telefone_comercial')
        fax_comercial = request.form.get('fax_comercial')
        email = request.form.get('email')
        bip = request.form.get('bip')
        data_nascimento = from_str_to_datetime_or_none(request.form.get('data_nascimento'))
        mes = to_int_or_none( request.form.get('mes') )

        if data_nascimento:
            mes = data_nascimento.month
        else:
            mes = 0

        #Criar dicionário com os dados
        dicionario = {
            "nome":nome,
            "rg":rg,
            "cpf":cpf,
            "endereco":endereco,
            "complemento":complemento,
            "bairro":bairro,
            "cidade":cidade,
            "cep":cep,
            "estado":estado,
            "pais":pais,
            "telefone":telefone,
            "fax":fax,
            "celular":celular,
            "telefone_comercial":telefone_comercial,
            "fax_comercial":fax_comercial,
            "email":email,
            "bip":bip,
            "data_nascimento":data_nascimento,
            "mes":mes
        }
        if pk:
            dicionario['id'] = pk
        cliente = Cliente(**dicionario)
        mensagem = None
        try:
            contexto['tipo_mensagem'] = 'success'
            if pk:
                db.session.merge(cliente)
            else:
                db.session.add(cliente)
            db.session.commit()
            id_cadastro = cliente.id
            if pk:
                flash( u'Cliente {0} atualizado com sucesso.'.format(id_cadastro), 'success')
            else:
                flash( u'Cliente {0} cadastrado com sucesso.'.format(id_cadastro), 'success')
            return redirect(url_for('cliente.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar cliente.'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Cliente.query.filter_by(id=pk).one()
        contexto['model'] = Cliente.to_dict(data, cliente_colunas)
    return render_template('cliente/cadastro.html', **contexto)


@cliente_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Cliente.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404


def get_filter(_nome, _telefone, _celular):
    lista_filtros = []
    if _nome:
        lista_filtros.append( Cliente.nome.like('%'+_nome+'%') )
    if _telefone:
        lista_filtros.append( Cliente.telefone.like('%'+_telefone+'%') )
    if _celular:
        lista_filtros.append( Cliente.celular.like('%'+_celular+'%') )
    return or_( *lista_filtros )


@cliente_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _sort_order = request.args.get('sort_order', '')
    _sort_direction = request.args.get('sort_direction', 'asc')

    _nome = request.args.get('nome', '')
    _telefone = request.args.get('telefone', '')
    _celular = request.args.get('celular', '')
    _limit = _offset + _limit
    items = []
    filtro = get_filter(_nome, _telefone, _celular)
    try:
        fetch = Cliente.query.filter( filtro )
        fetch = Cliente.sorting_data(fetch, _sort_order, _sort_direction)
        fetch = fetch.slice(_offset, _limit).all()
        colunas = [ col.name for col in Cliente.__table__._columns ]
        for dado in fetch:
            items.append( Cliente.to_dict(dado, cliente_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@cliente_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _nome = request.args.get('nome', '')
    _telefone = request.args.get('telefone', '')
    _celular = request.args.get('celular', '')
    count = 0
    filtro = get_filter(_nome, _telefone, _celular)
    try:
        count = Cliente.query.filter( filtro ).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@cliente_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Cliente.query.filter_by(id=pk).one_or_none()
    if data is not None:
        return Response(response=json.dumps( Cliente.to_dict(data, cliente_colunas) ), status=200, mimetype="application/json")
    return '',404
