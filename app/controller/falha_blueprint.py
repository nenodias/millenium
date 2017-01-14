# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (Blueprint, render_template, request, redirect, url_for, flash, 
    jsonify, render_template, Response)
from app import auth_require
from app import db
from app.models import Falha

falha_blueprint = Blueprint('falha', __name__)

falha_colunas = [ col.name for col in Falha.__table__._columns ]

@falha_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _descricao = request.args.get('descricao', '')
    contexto['model'] = {
        'descricao':_descricao,
    }
    contexto['inherit']='layout.html'
    return render_template('falha/consulta.html', **contexto)

@falha_blueprint.route('/index/ajax')
@auth_require()
def index_ajax():
    contexto = {}
    _descricao = request.args.get('descricao', '')
    contexto['model'] = {
        'descricao':_descricao,
    }
    contexto['inherit']='ajax.html'
    return render_template('falha/consulta.html', **contexto)

@falha_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@falha_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        descricao = request.form.get("descricao")
      
        #Criar dicionário com os dados
        dicionario = {
            "descricao":descricao,
        }
        if pk:
            dicionario['id'] = pk
        modelo = Falha(**dicionario)
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
                flash( u'Falha {0} atualizada com sucesso'.format(id_cadastro), 'success')
            else:
                flash( u'Falha {0} cadastrada com sucesso'.format(id_cadastro), 'success')
            return redirect(url_for('falha.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar falha'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Falha.query.filter_by(id=pk).one()
        contexto['model'] = Falha.to_dict(data, falha_colunas)
    return render_template('falha/cadastro.html', **contexto)


@falha_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Falha.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404

@falha_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _descricao = request.args.get('descricao', '')
    _limit = _offset + _limit
    items = []

    try:
        fetch = Falha.query.filter(Falha.descricao.like('%'+_descricao+'%')).slice(_offset, _limit).all()
        colunas = [ col.name for col in Falha.__table__._columns ]
        for dado in fetch:
            items.append( Falha.to_dict(dado, falha_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@falha_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _descricao = request.args.get('descricao', '')
    count = 0
    try:
        count = Falha.query.filter(Falha.descricao.like('%'+_descricao+'%')).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@falha_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Falha.query.filter_by(id=pk).one()
    if data:
        return Response(response=json.dumps( Falha.to_dict(data, falha_colunas) ), status=200, mimetype="application/json")
    return '',404
