# -*- coding: utf-8 -*-
import json
import logging
from flask import (
    Blueprint,
    request,
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
from app.models import Historico, or_, and_
from app.schema import historico_schema, historicos_schema
from .rest_util import PaginateRequest

api = Blueprint('rest_historico', __name__)

@api.route('/count', methods = ['get'])
@auth_require()
def count():
    page = PaginateRequest(request, Historico, historico_schema, historicos_schema)
    _numero_ordem = to_int_or_none( request.args.get("numero_ordem") )
    _id_cliente = to_int_or_none( request.args.get('id_cliente', '') )
    _id_veiculo = to_int_or_none( request.args.get('id_veiculo', '') )
    _id_tecnico = to_int_or_none( request.args.get('id_tecnico', '') )
    _data = from_str_to_date_or_none( request.args.get('data', '') )
    _tipo = request.args.get('tipo', '')
    filtro = get_filter(_numero_ordem, _id_cliente, _id_veiculo, _id_tecnico, _data, _tipo)
    return page.query_count(filtro)

@api.route('/', methods = ['get'])
@auth_require()
def list():
    page = PaginateRequest(request, Historico, historico_schema, historicos_schema)
    _numero_ordem = to_int_or_none( request.args.get("numero_ordem") )
    _id_cliente = to_int_or_none( request.args.get('id_cliente', '') )
    _id_veiculo = to_int_or_none( request.args.get('id_veiculo', '') )
    _id_tecnico = to_int_or_none( request.args.get('id_tecnico', '') )
    _data = from_str_to_date_or_none( request.args.get('data', '') )
    _tipo = request.args.get('tipo', '')
    filtro = get_filter(_numero_ordem, _id_cliente, _id_veiculo, _id_tecnico, _data, _tipo)
    return page.query_fetch( filtro )


@api.route('/<pk>', methods = ['get'])
@auth_require()
def get(pk):
    page = PaginateRequest(request, Historico, historico_schema, historicos_schema)
    data = page.query_one(pk)
    if data:
        return data
    return page.response({"message": "Registro não encontrado"}, 404)

@api.route('/<pk>', methods = ['delete'])
@auth_require()
def delete(pk):
    page = PaginateRequest(request, Historico, historico_schema, historicos_schema)
    return page.delete_one(pk, db)


@api.route('/', defaults={'pk':None}, methods = ['post'])
@api.route('/<pk>', methods = ['post'])
@auth_require()
def post(pk):
    page = PaginateRequest(request, Historico, historico_schema, historicos_schema)
    return page.post(pk, db)


def get_filter(_numero_ordem, _id_cliente, _id_veiculo, _id_tecnico, _data, _tipo):
    lista_filtros = []
    if _numero_ordem:
        lista_filtros.append( Historico.numero_ordem==_numero_ordem)
    if _id_cliente:
        lista_filtros.append( Historico.id_cliente==_id_cliente)
    if _id_veiculo:
        lista_filtros.append( Historico.id_veiculo==_id_veiculo )
    if _id_tecnico:
        lista_filtros.append( Historico.id_tecnico==_id_tecnico )
    if _data:
        _end = final_date_day(_data)
        lista_filtros.append( Historico.data.between(_data, _end) )
    if _tipo:
        lista_filtros.append( Historico.tipo==_tipo )
    return and_( *lista_filtros )
