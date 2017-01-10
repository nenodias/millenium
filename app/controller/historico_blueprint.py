# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (Blueprint, render_template, request, redirect, url_for, flash, 
    jsonify, render_template, Response)
from app import auth_require, db
from app.utils import to_int_or_none, from_str_to_datetime_or_none, from_str_to_date_or_none, final_date_day
from app.models import Historico, HistoricoItem, or_, and_

historico_blueprint = Blueprint('historico', __name__)

tupla_tipo_historico = ( ('O.S.', 'Ordem de Serviço'), ('Orç.', 'Orçamento') )
tupla_tipo_item = ( ('S', 'Serviço'),('F', 'Falha'), ('P', 'Peça') )

items_colunas = ['id','ordem','tipo','descricao','quantidade','valor']
historico_colunas = [ 'id', 'id_cliente', 'id_veiculo', 'id_tecnico', 'numero_ordem', 'placa', 'sistema', 'data', 'tipo', 'valor_total', 'observacao', ('items', items_colunas )]

@historico_blueprint.route('/')
@auth_require()
def index():
    contexto = {}
    _id_cliente = request.args.get('id_cliente', '')
    _id_veiculo = request.args.get('id_veiculo', '')
    _id_tecnico = request.args.get('id_tecnico', '')
    _data = request.args.get('data', '')
    _tipo = request.args.get('tipo', '')
    contexto['model'] = {
        'id_cliente':_id_cliente,
        'id_veiculo': _id_veiculo,
        'id_tecnico':_id_tecnico,
        'data':_data,
        'tipo':_tipo,
    }
    contexto['tupla_tipo_historico'] = tupla_tipo_historico
    return render_template('historico/consulta.html', **contexto)

@historico_blueprint.route('/form/', defaults={'pk':None}, methods = ['post', 'get'])
@historico_blueprint.route('/form/<pk>', methods = ['post', 'get'])
@auth_require()
def form(pk):
    #Pega os dados dos campos na tela
    contexto = {}
    contexto['model'] = {}
    if request.method == 'POST':
        
        id_cliente = to_int_or_none( request.form.get("id_cliente") )
        id_veiculo = to_int_or_none( request.form.get("id_veiculo") )
        id_tecnico = to_int_or_none( request.form.get("id_tecnico") )
        numero_ordem = to_int_or_none( request.form.get("numero_ordem") )
        placa = request.form.get("placa")
        sistema = to_int_or_none( request.form.get("sistema") )
        data = from_str_to_datetime_or_none( request.form.get("data") )
        tipo = request.form.get("tipo")
        valor_total = float( request.form.get("valor_total") )
        observacao = request.form.get("observacao")
     
        #Criar dicionário com os dados
        dicionario = {
            'id_cliente':id_cliente,
            'id_veiculo':id_veiculo,
            'id_tecnico':id_tecnico,
            'numero_ordem':numero_ordem,
            'placa':placa,
            'sistema':sistema,
            'data':data,
            'tipo':tipo,
            'valor_total':valor_total,
            'observacao':observacao,
        }
        if pk:
            dicionario['id'] = pk
        historico = Historico(**dicionario)

        items = to_int_or_none( request.form.get("items") )
        list_items = []
        if items:
            for i in range(items):
                id = to_int_or_none( request.form.get('item_id_'+i) )
                ordem = to_int_or_none( request.form.get('item_ordem_'+i) )
                tipo = request.form.get('item_tipo_'+i)
                descricao = request.form.get('item_descricao_'+i)
                quantidade = to_int_or_none( request.form.get('item_quantidade_'+i) )
                valor = float( request.form.get('item_valor_'+i) )
                item_dict = { 
                        'id':id,
                        'ordem':ordem,
                        'tipo':tipo,
                        'descricao':descricao,
                        'quantidade':quantidade,
                        'valor':valor,
                        'id_historico': historico.id
                        }
                list_items.append( HistoricoItem(**item_dict) )

        mensagem = None
        try:
            contexto['tipo_mensagem'] = 'success'
            if pk:
                db.session.merge(historico)
            else:
                db.session.add(historico)
            if list_items:
                for item in list_items:
                    if item.id:
                        db.sessopn.merge(item)
                    else:
                        db.sessopn.add(item)
            db.session.commit()
            id_cadastro = historico.id
            if pk:
                flash( u'Historico {0} atualizado com sucesso'.format(id_cadastro), 'success')
            else:
                flash( u'Historico {0} cadastrado com sucesso'.format(id_cadastro), 'success')
            return redirect(url_for('historico.index'))
        except Exception as ex:
            print(ex)
            contexto['mensagem'] = u'Erro ao cadastrar historico'
            contexto['tipo_mensagem'] = 'danger'
    elif pk:
        data = Historico.query.filter_by(id=pk).one()
        contexto['model'] = Historico.to_dict(data, historico_colunas)
    return render_template('historico/cadastro.html', **contexto)


@historico_blueprint.route('/delete/<pk>', methods = ['post'])
@auth_require()
def delete(pk):
    data = Historico.query.filter_by(id=pk).one()
    if data:
        try:
            db.session.delete(data)
            db.session.commit()
            return '', 200
        except Exception as ex:
            print(ex)
    return '',404


def get_filter(_id_cliente, _id_veiculo, _id_tecnico, _data, _tipo):
    lista_filtros = []
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


@historico_blueprint.route('/ajax', methods = ['get'])
@auth_require()
def ajax():
    _limit = int(request.args.get('limit','10'))
    _offset = int(request.args.get('offset','0'))

    _id_cliente = to_int_or_none( request.args.get('id_cliente', '') )
    _id_veiculo = to_int_or_none( request.args.get('id_veiculo', '') )
    _id_tecnico = to_int_or_none( request.args.get('id_tecnico', '') )
    _data = from_str_to_date_or_none( request.args.get('data', '') )
    _tipo = request.args.get('tipo', '')

    _limit = _offset + _limit
    items = []
    filtro = get_filter(_id_cliente, _id_veiculo, _id_tecnico, _data, _tipo)
    try:
        fetch = Historico.query.filter( filtro ).slice(_offset, _limit).all()
        colunas = [ col.name for col in Historico.__table__._columns ]
        for dado in fetch:
            items.append( Historico.to_dict(dado, historico_colunas) )
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( items ), status=200, mimetype="application/json")

@historico_blueprint.route('/count', methods = ['get'])
@auth_require()
def count():
    
    _id_cliente = to_int_or_none( request.args.get('id_cliente', '') )
    _id_veiculo = to_int_or_none( request.args.get('id_veiculo', '') )
    _id_tecnico = to_int_or_none( request.args.get('id_tecnico', '') )
    _data = from_str_to_date_or_none( request.args.get('data', '') )
    _tipo = request.args.get('tipo', '')

    count = 0
    filtro = get_filter(_id_cliente, _id_veiculo, _id_tecnico, _data, _tipo)
    try:
        count = Historico.query.filter( filtro ).count()
    except Exception as ex:
        print(ex)
    return Response(response=json.dumps( {"count":count} ), status=200, mimetype="application/json")

@historico_blueprint.route('/ajax/<pk>', methods = ['get'])
@auth_require()
def ajax_by_id(pk):
    data = Historico.query.filter_by(id=pk).one()
    if data:
        return Response(response=json.dumps( Historico.to_dict(data, historico_colunas) ), status=200, mimetype="application/json")
    return '',404
