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
from app import db
from app.utils import to_int_or_none, from_str_to_datetime_or_none
from app.models import Veiculo, or_, and_, desc, veiculo_colunas

veiculo_blueprint = Blueprint('veiculo', __name__)


@veiculo_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _id_cliente = request.args.get('id_cliente', '')
    _id_modelo = request.args.get('id_modelo', '')
    _placa = request.args.get('placa', '')
    contexto['model'] = {
        'id_cliente':_id_cliente,
        'id_modelo': _id_modelo,
        'placa':_placa
    }
    return render_template('veiculo/consulta.html', **contexto)

@veiculo_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@veiculo_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        id_cliente = to_int_or_none(request.form.get('id_cliente'))
        id_modelo = to_int_or_none(request.form.get('id_modelo'))
        placa = request.form.get('placa')
        pais = request.form.get('pais')
        cor = request.form.get('cor')
        combustivel = request.form.get('combustivel')
        renavam = request.form.get('renavam')
        chassi = request.form.get('chassi')
        ano = request.form.get('ano')

        #Criar dicionário com os dados
        dicionario = {
            'id_cliente':id_cliente,
            'id_modelo':id_modelo,
            'placa':placa,
            'pais':pais,
            'cor':cor,
            'combustivel':combustivel,
            'renavam':renavam,
            'chassi':chassi,
            'ano':ano,
        }
        if pk:
            dicionario['id'] = pk
        veiculo = Veiculo(**dicionario)
        mensagem = None
        try:
            contexto['tipo_mensagem'] = 'success'
            if pk:
                db.session.merge(veiculo)
            else:
                db.session.add(veiculo)
            db.session.commit()
            id_cadastro = veiculo.id
            if pk:
                flash( u'Veículo {0} atualizado com sucesso.'.format(id_cadastro), 'success')
            else:
                flash( u'Veículo {0} cadastrado com sucesso.'.format(id_cadastro), 'success')
            return redirect(url_for('veiculo.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar veículo.'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Veiculo.query.filter_by(id=pk).one()
        contexto['model'] = Veiculo.to_dict(data, veiculo_colunas)
    return render_template('veiculo/cadastro.html', **contexto)


@veiculo_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Veiculo.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404


def get_filter(_id_cliente, _id_modelo, _placa):
    lista_filtros = []
    if _id_cliente:
        lista_filtros.append( Veiculo.id_cliente==_id_cliente)
    if _id_modelo:
        lista_filtros.append( Veiculo.id_modelo==_id_modelo )
    if _placa:
        lista_filtros.append( Veiculo.placa.like('%'+_placa+'%') )
    return and_( *lista_filtros )


@veiculo_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _sort_order = request.args.get('sort_order', '')
    _sort_direction = request.args.get('sort_direction', 'asc')

    _id_cliente = to_int_or_none(request.args.get('id_cliente', ''))
    _id_modelo = to_int_or_none(request.args.get('id_modelo', ''))
    _placa = request.args.get('placa', '')
    _limit = _offset + _limit
    items = []
    filtro = get_filter(_id_cliente, _id_modelo, _placa)
    try:
        fetch = Veiculo.query.filter( filtro )
        fetch = Veiculo.sorting_data(fetch, _sort_order, _sort_direction)
        fetch = fetch.slice(_offset, _limit).all()
        colunas = [ col.name for col in Veiculo.__table__._columns ]
        for dado in fetch:
            items.append( Veiculo.to_dict(dado, veiculo_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@veiculo_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _id_cliente = to_int_or_none(request.args.get('id_cliente', ''))
    _id_modelo = to_int_or_none(request.args.get('id_modelo', ''))
    _placa = request.args.get('placa', '')
    count = 0
    filtro = get_filter(_id_cliente, _id_modelo, _placa)
    try:
        count = Veiculo.query.filter( filtro ).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@veiculo_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Veiculo.query.filter_by(id=pk).one_or_none()
    if data is not None:
        return Response(response=json.dumps( Veiculo.to_dict(data, veiculo_colunas) ), status=200, mimetype="application/json")
    return '',404
