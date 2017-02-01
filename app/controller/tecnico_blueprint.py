# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (Blueprint, render_template, request, redirect, url_for, flash, 
    jsonify, render_template, Response)
from app import auth_require
from app import db
from app.models import Tecnico, desc

tecnico_blueprint = Blueprint('tecnico', __name__)

tecnico_colunas = [ 'id', 'nome' ]

@tecnico_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _nome = request.args.get('nome', '')
    contexto['model'] = {
        'nome':_nome,
    }
    return render_template('tecnico/consulta.html', **contexto)

@tecnico_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@tecnico_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        nome = request.form.get("nome")
      
        #Criar dicionário com os dados
        dicionario = {
            "nome":nome,
        }
        if pk:
            dicionario['id'] = pk
        modelo = Tecnico(**dicionario)
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
                flash( u'Técnico {0} atualizado com sucesso.'.format(id_cadastro), 'success')
            else:
                flash( u'Técnico {0} cadastrado com sucesso.'.format(id_cadastro), 'success')
            return redirect(url_for('tecnico.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar tecnico.'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Tecnico.query.filter_by(id=pk).one()
        contexto['model'] = Tecnico.to_dict(data, tecnico_colunas)
    return render_template('tecnico/cadastro.html', **contexto)


@tecnico_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Tecnico.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404

@tecnico_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _sort_order = request.args.get('sort_order', '')
    _sort_direction = request.args.get('sort_direction', 'asc')
    
    _nome = request.args.get('nome', '')
    _limit = _offset + _limit
    items = []

    try:
        filtro = Tecnico.nome.like('%'+_nome+'%')
        fetch = Tecnico.query.filter( filtro )
        fetch = Tecnico.sorting_data(fetch, _sort_order, _sort_direction)
        fetch = fetch.slice(_offset, _limit).all()
        colunas = [ col.name for col in Tecnico.__table__._columns ]
        for dado in fetch:
            items.append( Tecnico.to_dict(dado, tecnico_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@tecnico_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _nome = request.args.get('nome', '')
    count = 0
    try:
        count = Tecnico.query.filter(Tecnico.nome.like('%'+_nome+'%')).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@tecnico_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Tecnico.query.filter_by(id=pk).one()
    if data:
        return Response(response=json.dumps( Tecnico.to_dict(data, tecnico_colunas) ), status=200, mimetype="application/json")
    return '',404
