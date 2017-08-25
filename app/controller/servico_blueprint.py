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
from app import auth_require, db
from app.models import Servico, desc

servico_blueprint = Blueprint('servico', __name__)

servico_colunas = [col.name for col in Servico.__table__._columns]


@servico_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _descricao = request.args.get('descricao', '')
    contexto['model'] = {
        'descricao':_descricao,
    }
    contexto['inherit']='layout.html'
    return render_template('servico/consulta.html', **contexto)

@servico_blueprint.route('/index/ajax')
@auth_require()
def index_ajax():
    contexto = {}
    _descricao = request.args.get('descricao', '')
    contexto['model'] = {
        'descricao':_descricao,
    }
    contexto['inherit']='ajax.html'
    return render_template('servico/consulta.html', **contexto)

@servico_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@servico_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        descricao = request.form.get("descricao")
        valor = float(request.form.get("valor"))

        #Criar dicionário com os dados
        dicionario = {
            "descricao":descricao,
            "valor":valor
        }
        if pk:
            dicionario['id'] = pk
        modelo = Servico(**dicionario)
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
                flash( u'Serviço {0} atualizado com sucesso.'.format(id_cadastro), 'success')
            else:
                flash( u'Serviço {0} cadastrado com sucesso.'.format(id_cadastro), 'success')
            return redirect(url_for('servico.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar serviço.'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Servico.query.filter_by(id=pk).one()
        contexto['model'] = Servico.to_dict(data, servico_colunas)
    return render_template('servico/cadastro.html', **contexto)


@servico_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Servico.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404

@servico_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))
    _sort_order = request.args.get('sort_order', '')
    _sort_direction = request.args.get('sort_direction', 'asc')

    _descricao = request.args.get('descricao', '')
    _limit = _offset + _limit
    items = []

    try:
        filtro = Servico.descricao.like('%'+_descricao+'%')
        fetch = Servico.query.filter( filtro )
        fetch = Servico.sorting_data(fetch, _sort_order, _sort_direction)
        fetch = fetch.slice(_offset, _limit).all()
        colunas = [ col.name for col in Servico.__table__._columns ]
        for dado in fetch:
            items.append( Servico.to_dict(dado, servico_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@servico_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    _descricao = request.args.get('descricao', '')
    count = 0
    try:
        count = Servico.query.filter(Servico.descricao.like('%'+_descricao+'%')).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@servico_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Servico.query.filter_by(id=pk).one()
    if data:
        return Response(response=json.dumps( Servico.to_dict(data, servico_colunas) ), status=200, mimetype="application/json")
    return '',404
