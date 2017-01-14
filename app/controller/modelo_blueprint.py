# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (Blueprint, render_template, request, redirect, url_for, flash, 
    jsonify, render_template, Response)
from app import auth_require, db
from app.utils import to_int_or_none
from app.models import Modelo, or_, and_, modelo_colunas

modelo_blueprint = Blueprint('modelo', __name__)

@modelo_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _nome = request.args.get('nome', '')
    _id_monta = request.args.get('id_monta', '')
    contexto['model'] = {
        'nome':_nome,
        'id_monta': _id_monta
    }
    return render_template('modelo/consulta.html', **contexto)

@modelo_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@modelo_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        nome = request.form.get("nome")
        codvei_ea = to_int_or_none(request.form.get("codvei_ea") )
        id_monta = int(request.form.get("id_monta") )
      
        #Criar dicionário com os dados
        dicionario = {
            "nome":nome,
            "codvei_ea":codvei_ea,
            "id_monta":id_monta
        }
        if pk:
            dicionario['id'] = pk
        modelo = Modelo(**dicionario)
        mensagem = None
        try:
            contexto['tipo_mensagem'] = 'success'
            if pk:
                db.session.merge(modelo)
            else:
                db.session.add(modelo)
            db.session.commit()
            id_cadastro = modelo.id
            if pk:
                flash( u'Modelo {0} atualizado com sucesso'.format(id_cadastro), 'success')
            else:
                flash( u'Modelo {0} cadastrado com sucesso'.format(id_cadastro), 'success')
            return redirect(url_for('modelo.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar modelo'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Modelo.query.filter_by(id=pk).one()
        contexto['model'] = Modelo.to_dict(data, modelo_colunas)
    return render_template('modelo/cadastro.html', **contexto)


@modelo_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Modelo.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404


def get_filter(_nome, _id_monta):
    lista_filtros = []
    if _nome:
        lista_filtros.append( Modelo.nome.like('%'+_nome+'%') )
    if _id_monta:
        lista_filtros.append( Modelo.id_monta==_id_monta  )
    return and_( *lista_filtros )


@modelo_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _nome = request.args.get('nome', '')
    _id_monta = to_int_or_none(request.args.get('id_monta'))
    _limit = _offset + _limit
    items = []
    filtro = get_filter(_nome, _id_monta)
    try:
        fetch = Modelo.query.filter( filtro ).slice(_offset, _limit).all()
        colunas = [ col.name for col in Modelo.__table__._columns ]
        for dado in fetch:
            items.append( Modelo.to_dict(dado, modelo_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@modelo_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _nome = request.args.get('nome', '')
    _id_monta = to_int_or_none(request.args.get('id_monta'))
    count = 0
    filtro = get_filter(_nome, _id_monta)
    try:
        count = Modelo.query.filter( filtro ).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@modelo_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Modelo.query.filter_by(id=pk).one()
    if data:
        return Response(response=json.dumps( Modelo.to_dict(data, modelo_colunas) ), status=200, mimetype="application/json")
    return '',404
