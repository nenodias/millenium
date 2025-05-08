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
from app.utils import (
    to_int_or_none,
    from_str_to_datetime_or_none,
    from_str_to_date_or_none,
    final_date_day,
    to_float_or_zero
)
from app.models import (Lembrete, lembrete_colunas, or_, and_)

lembrete_blueprint = Blueprint('lembrete', __name__)


@lembrete_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _numero_ordem = request.args.get("numero_ordem",'')
    _id_cliente = request.args.get('id_cliente', '')
    _id_veiculo = request.args.get('id_veiculo', '')
    _data_notificacao = request.args.get('data_notificacao', '')
    contexto['model'] = {
        'id_cliente':_id_cliente,
        'id_veiculo': _id_veiculo,
        'data_notificacao':_data_notificacao,
    }
    return render_template('lembrete/consulta.html', **contexto)

@lembrete_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@lembrete_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {
    }
    if request.method == 'POST':

        id_cliente = to_int_or_none( request.form.get("id_cliente") )
        id_veiculo = to_int_or_none( request.form.get("id_veiculo") )
        data_notificacao = from_str_to_datetime_or_none( request.form.get("data_notificacao") )
        texto = request.form.get("texto")

        #Criar dicionário com os dados
        dicionario = {
            'id_cliente':id_cliente,
            'id_veiculo':id_veiculo,
            'data_notificacao':data_notificacao,
            'texto':texto,
        }
        if pk:
            dicionario['id'] = pk
        lembrete = Lembrete(**dicionario)
        if not lembrete.id:
            db.session.add(lembrete)
            db.session.flush()
            db.session.refresh(lembrete)
        mensagem = None
        try:
            db.session.commit()
            id_cadastro = lembrete.id
            if pk:
                flash( 'Lembrete {0} atualizado com sucesso.'.format(id_cadastro), 'success')
            else:
                flash( 'Lembrete {0} cadastrado com sucesso.'.format(id_cadastro), 'success')
            return redirect(url_for('lembrete.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar lembrete.'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Lembrete.query.filter_by(id=pk).one()
        contexto['model'] = Lembrete.to_dict(data, lembrete_colunas)
    return render_template('lembrete/cadastro.html', **contexto)


@lembrete_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Lembrete.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404


def get_filter(_id_cliente, _id_veiculo,_data):
    lista_filtros = []
    if _id_cliente:
        lista_filtros.append( Lembrete.id_cliente==_id_cliente)
    if _id_veiculo:
        lista_filtros.append( Lembrete.id_veiculo==_id_veiculo )
    if _data:
        _end = final_date_day(_data)
        lista_filtros.append( Lembrete.data_notificacao.between(_data, _end) )
    return and_( *lista_filtros )


@lembrete_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _sort_order = request.args.get('sort_order', '')
    _sort_direction = request.args.get('sort_direction', 'asc')

    _id_cliente = to_int_or_none( request.args.get('id_cliente', '') )
    _id_veiculo = to_int_or_none( request.args.get('id_veiculo', '') )
    _data_notificacao = from_str_to_date_or_none( request.args.get('data_notificacao', '') )

    _limit = _offset + _limit
    items = []
    filtro = get_filter(_id_cliente, _id_veiculo,_data_notificacao)
    try:
        fetch = Lembrete.query.filter( filtro )
        fetch = Lembrete.sorting_data(fetch, _sort_order, _sort_direction)
        fetch = fetch.slice(_offset, _limit).all()
        for dado in fetch:
            items.append( Lembrete.to_dict(dado, lembrete_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@lembrete_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _id_cliente = to_int_or_none( request.args.get('id_cliente', '') )
    _id_veiculo = to_int_or_none( request.args.get('id_veiculo', '') )
    _data_notificacao = from_str_to_date_or_none( request.args.get('data_notificacao', '') )

    count = 0
    filtro = get_filter(_id_cliente, _id_veiculo,_data_notificacao)
    try:
        count = Lembrete.query.filter( filtro ).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@lembrete_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Lembrete.query.filter_by(id=pk).one_or_none()
    if data is not None:
        return Response(response=json.dumps( Lembrete.to_dict(data, lembrete_colunas) ), status=200, mimetype="application/json")
    return '',404
