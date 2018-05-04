# -*- coding: utf-8 -*-
import json
import logging
from flask import (
    Blueprint,
    request,
)
from app.authentication import auth_require
from app import db
from app.models import Cliente, or_
from app.schema import cliente_schema, clientes_schema
from .rest_util import PaginateRequest

api = Blueprint('rest_client', __name__)

@api.route('/count', methods = ['get'])
@auth_require()
def count():
    page = PaginateRequest(request, Cliente, cliente_schema, clientes_schema)
    _nome = request.args.get('nome', None)
    _telefone = request.args.get('telefone', None)
    _celular = request.args.get('celular', None)
    filtro = get_filter(_nome, _telefone, _celular)
    return page.query_count(filtro)

@api.route('/', methods = ['get'])
@auth_require()
def list():
    page = PaginateRequest(request, Cliente, cliente_schema, clientes_schema)
    _nome = request.args.get('nome', None)
    _telefone = request.args.get('telefone', None)
    _celular = request.args.get('celular', None)
    filtro = get_filter(_nome, _telefone, _celular)
    return page.query_fetch( filtro )


@api.route('/<pk>', methods = ['get'])
@auth_require()
def get(pk):
    page = PaginateRequest(request, Cliente, cliente_schema, clientes_schema)
    data = page.query_one(pk)
    if data:
        return data
    return page.response({"message": "Registro não encontrado"}, 404)

@api.route('/<pk>', methods = ['delete'])
@auth_require()
def delete(pk):
    page = PaginateRequest(request, Cliente, cliente_schema, clientes_schema)
    return page.delete_one(pk, db)


@api.route('/', defaults={'pk':None}, methods = ['post'])
@api.route('/<pk>', methods = ['post'])
@auth_require()
def post(pk):
    page = PaginateRequest(request, Cliente, cliente_schema, clientes_schema)
    return page.post(pk, db)


def get_filter(_nome, _telefone, _celular):
    filters = []
    if _nome:
        filters.append( Cliente.nome.like('%'+_nome+'%') )
    if _telefone:
        filters.append( Cliente.telefone.like('%'+_telefone+'%') )
    if _celular:
        filters.append( Cliente.celular.like('%'+_celular+'%') )
    return or_( *filters )
