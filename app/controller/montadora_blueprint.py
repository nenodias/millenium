# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (Blueprint, render_template, request, redirect, url_for, flash, 
    jsonify, render_template, Response)
from app import auth_require
from app import db
from app.models import Montadora, tupla_origem, montadora_colunas

montadora_blueprint = Blueprint('montadora', __name__)

@montadora_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _nome = request.args.get('nome', '')
    contexto['model'] = {
        'nome':_nome,
    }
    return render_template('montadora/consulta.html', **contexto)

@montadora_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@montadora_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        nome = request.form.get("nome")
        origem = request.form.get("origem")
        codmon_ea = int(request.form.get("codmon_ea"))
      
        #Criar dicionário com os dados
        dicionario = {
            "nome":nome,
            "origem":origem,
            "codmon_ea":codmon_ea
        }
        if pk:
            dicionario['id'] = pk
        modelo = Montadora(**dicionario)
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
                flash( u'Montadora {0} atualizada com sucesso'.format(id_cadastro), 'success')
            else:
                flash( u'Montadora {0} cadastrada com sucesso'.format(id_cadastro), 'success')
            return redirect(url_for('montadora.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar montadora'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Montadora.query.filter_by(id=pk).one()
        contexto['model'] = Montadora.to_dict(data, montadora_colunas)
    return render_template('montadora/cadastro.html', **contexto)


@montadora_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Montadora.query.filter_by(id=pk).one()
    if data:
        if  len(data.modelos) > 0:
            lista_ids = []
            for modelo in data.modelos:
                lista_ids.append(modelo.id)
            return 'Existem os seguintes modelos cadastrados para essa montadora %s'%(lista_ids), 500
        else:
            try:
                db.session.delete(data)
                db.session.commit()
                return '', 200
            except Exception as ex:
                print(ex)
    return '',404

@montadora_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _nome = request.args.get('nome', '')
    _limit = _offset + _limit
    items = []

    try:
        fetch = Montadora.query.filter(Montadora.nome.like('%'+_nome+'%')).slice(_offset, _limit).all()
        for dado in fetch:
            items.append( Montadora.to_dict(dado, montadora_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@montadora_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _nome = request.args.get('nome', '')
    count = 0
    try:
        count = Montadora.query.filter(Montadora.nome.like('%'+_nome+'%')).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@montadora_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Montadora.query.filter_by(id=pk).one()
    if data:
        return Response(response=json.dumps( Montadora.to_dict(data, montadora_colunas) ), status=200, mimetype="application/json")
    return '',404
