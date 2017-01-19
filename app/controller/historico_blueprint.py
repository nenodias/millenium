# -*- coding: utf-8 -*-
import json
from pdb import set_trace
from flask import (Blueprint, render_template, request, redirect, url_for, flash, send_file,
    jsonify, render_template, Response)
from app import auth_require, db
from app.utils import (to_int_or_none, from_str_to_datetime_or_none, from_str_to_date_or_none,
 final_date_day, to_float_or_zero)
from app.models import (Historico, HistoricoItem, Vistoria, Cliente, Veiculo, Modelo, Montadora, or_, and_,
    tupla_tipo_historico, tupla_tipo_item, items_colunas, historico_colunas, cliente_colunas, 
    veiculo_colunas, modelo_colunas, montadora_colunas)

historico_blueprint = Blueprint('historico', __name__)

def get_tipo(tipo):
    if tipo == 'F':
        return 'falha'
    elif tipo == 'S':
        return 'servico'
    else:
        return 'peca'

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
    contexto['model'] = {
        "items":[],
        "vistoria":{}
    }
    contexto['tupla_tipo_historico'] = tupla_tipo_historico
    contexto['get_tipo'] = get_tipo
    contexto['tipo_servico'] = tupla_tipo_item[0][0]
    contexto['tipo_falha'] = tupla_tipo_item[1][0]
    contexto['tipo_peca'] = tupla_tipo_item[2][0]
    if request.method == 'POST':
        
        id_cliente = to_int_or_none( request.form.get("id_cliente") )
        id_veiculo = to_int_or_none( request.form.get("id_veiculo") )
        id_tecnico = to_int_or_none( request.form.get("id_tecnico") )
        numero_ordem = to_int_or_none( request.form.get("numero_ordem") )
        sistema = to_int_or_none( request.form.get("sistema") )
        data = from_str_to_datetime_or_none( request.form.get("data") )
        tipo = request.form.get("tipo")
        valor_total = to_float_or_zero( request.form.get("valor_total") )
        observacao = request.form.get("observacao")
        kilometragem = to_float_or_zero(request.form.get("kilometragem"))
        veiculo = Veiculo.query.filter_by(id=id_veiculo).one()
        if veiculo:
            placa = veiculo.placa
     
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
        if not historico.id:
            db.session.add(historico)
            db.session.flush()
            db.session.refresh(historico)
        if not numero_ordem:
            historico.numero_ordem = historico.id

        items = to_int_or_none( request.form.get("items") )
        list_items = []
        print('*'*10)
        print(historico.id)
        print('*'*10)
        if items:
            for i in range(items):
                idx = str(i)
                if request.form.get('item_descricao_'+ idx):
                    id = to_int_or_none( request.form.get('item_id_'+idx) )
                    ordem = to_int_or_none( request.form.get('item_ordem_'+idx) )
                    tipo = request.form.get('item_tipo_'+idx)
                    descricao = request.form.get('item_descricao_'+idx)
                    quantidade = to_int_or_none( request.form.get('item_quantidade_'+idx) )
                    valor = to_float_or_zero( request.form.get('item_valor_'+idx) )
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
            
            dados_vistoria = {
                'id':historico.id,
                'kilometragem':kilometragem,
                'observacao':''
            }
            vistoria = Vistoria(**dados_vistoria)
            try:
                db.session.merge(vistoria)
            except Exception as e:
                db.session.add(vistoria)
            contexto['tipo_mensagem'] = 'success'
            if pk:
                db.session.merge(historico)
            else:
                db.session.add(historico)
            pk_items = []
            if list_items:
                for item in list_items:
                    # salvando os items
                    item.ordem = len(pk_items)
                    if item.id:
                        db.session.merge(item)
                    else:
                        db.session.add(item)
                    pk_items.append( item.id )
                if pk_items:
                    # excluindo os items apagados que não foram passados na requisição
                    items_deletar = HistoricoItem.query.filter(
                        and_( 
                            ~HistoricoItem.id.in_( pk_items),
                            HistoricoItem.id_historico==historico.id 
                            )
                        ).all()
                    if items_deletar:
                        for item_delete in items_deletar:
                            db.session.delete(item_delete)
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


from app.relatorio import gerar_pdf

@historico_blueprint.route('/report/<pk>')
@auth_require()
def report(pk):
    data = Historico.query.filter_by(id=pk).one()
    dados = Historico.to_dict(data, historico_colunas)
    cliente = Cliente.query.filter_by(id=data.id_cliente).one()
    dados['cliente'] = Cliente.to_dict(cliente, cliente_colunas)
    veiculo = Veiculo.query.filter_by(id=data.id_cliente).one()
    dados['veiculo'] = Veiculo.to_dict(veiculo, veiculo_colunas)
    modelo = Modelo.query.filter_by(id=veiculo.id_modelo).one()
    dados['modelo'] = Modelo.to_dict(modelo, modelo_colunas)
    montadora = Montadora.query.filter_by(id=modelo.id_monta).one()
    dados['montadora'] = Modelo.to_dict(montadora, montadora_colunas)
    pdf_buffer = gerar_pdf( dados )
    return send_file(pdf_buffer, attachment_filename='relatorio.pdf',mimetype='application/pdf')