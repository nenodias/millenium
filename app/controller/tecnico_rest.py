# -*- coding: utf-8 -*-
import json
import logging
from flask import (
    Blueprint,
    request,
)
from app.authentication import auth_require
from app import db
from app.models import Tecnico
from app.schema import tecnico_schema, tecnicos_schema
from .rest_util import PaginateRequest

api = Blueprint('rest_tecnico', __name__)

@api.route('/count', methods = ['get'])
@auth_require()
def count():
    page = PaginateRequest(request, Tecnico, tecnico_schema, tecnicos_schema)
    _nome = request.args.get('nome', '')
    filtro = Tecnico.nome.like('%'+_nome+'%')
    return page.query_count(filtro)

@api.route('/', methods = ['get'])
@auth_require()
def list():
    page = PaginateRequest(request, Tecnico, tecnico_schema, tecnicos_schema)
    _nome = request.args.get('nome', '')
    filtro = Tecnico.nome.like('%'+_nome+'%')
    return page.query_fetch( filtro )


@api.route('/<pk>', methods = ['get'])
@auth_require()
def get(pk):
    page = PaginateRequest(request, Tecnico, tecnico_schema, tecnicos_schema)
    data = page.query_one(pk)
    if data:
        return data
    return page.response({"message": "Registro não encontrado"}, 404)

@api.route('/<pk>', methods = ['delete'])
@auth_require()
def delete(pk):
    page = PaginateRequest(request, Tecnico, tecnico_schema, tecnicos_schema)
    return page.delete_one(pk, db)


@api.route('/', defaults={'pk':None}, methods = ['post'])
@api.route('/<pk>', methods = ['post'])
@auth_require()
def post(pk):
    page = PaginateRequest(request, Tecnico, tecnico_schema, tecnicos_schema)
    return page.post(pk, db)