# -*- coding: utf-8 -*-
import pexpect
from flask import (
    Response, jsonify, request, render_template, redirect, session, url_for,
    current_app
    )
from app import (
    config, auth_require, app)
from app.utils import generate_hash
from .falha_blueprint import falha_blueprint
from .peca_blueprint import peca_blueprint
from .servico_blueprint import servico_blueprint
from .tecnico_blueprint import tecnico_blueprint
from .montadora_blueprint import montadora_blueprint
from .modelo_blueprint import modelo_blueprint
from .cliente_blueprint import cliente_blueprint
from .veiculo_blueprint import veiculo_blueprint
from .historico_blueprint import historico_blueprint
from .lembrete_blueprint import lembrete_blueprint


@app.route('/login', methods=['POST', 'GET'])
def login():
    contexto = {}
    if request.method == 'POST':
        usuario = None
        senha = None
        is_json = 'json' in request.content_type
        if is_json:
            dados = request.json
            usuario = dados['usuario']
            senha = dados['senha']
        else:
            usuario = request.form.get('usuario')
            senha = request.form.get('senha')
        def_user = config.DEFAULT_USERNAME
        def_pass = config.DEFAULT_PASSWORD
        user_valid = usuario == def_user
        pass_valid = senha == def_pass

        if is_json and user_valid and pass_valid:
            retorno = {"token": generate_hash(def_user, def_pass)}
            return jsonify(retorno)
        elif is_json:
            retorno = {"message": "Failed to login"}
            return jsonify(retorno)
        if user_valid and pass_valid:
            session['login'] = True
            return redirect(url_for('index'))
        else:
            contexto['tipo_mensagem'] = u'danger'
            contexto['mensagem'] = u'Usuário ou senha inválidos'

    return render_template('login.html', **contexto), 200


@app.route('/logout')
def logout():
    if 'login' in session:
        del session['login']
    return redirect(url_for('login'))


@app.route('/')
def index():
    return current_app.send_static_file('index.html'), 200


@app.route('/backup')
@auth_require()
def backup():
    uri = config.SQLALCHEMY_DATABASE_URI.split('://')[1]
    parte = uri.split('@')
    usuario, senha = parte[0].split(':')
    parte = parte[1].split(':')
    host = parte[0]
    database = None
    if '/' in parte[1]:
        porta, database = parte[1].split('/')
    else:
        porta = parte[1]
    # executando o pg_dump
    call = 'pg_dump -p '+porta+' -U '+usuario+' -h '+host+' -W'
    if not database:
        call += ' -C'
    ps = pexpect.spawn(call)
    ps.expect(':')
    ps.send('%s\n' % (senha))
    stdout = ps.read()
    return Response(stdout, content_type='text/plain; charset=utf-8')


app.register_blueprint(falha_blueprint, url_prefix='/falha')
app.register_blueprint(peca_blueprint, url_prefix='/peca')
app.register_blueprint(servico_blueprint, url_prefix='/servico')
app.register_blueprint(tecnico_blueprint, url_prefix='/tecnico')
app.register_blueprint(montadora_blueprint, url_prefix='/montadora')
app.register_blueprint(modelo_blueprint, url_prefix='/modelo')
app.register_blueprint(cliente_blueprint, url_prefix='/cliente')
app.register_blueprint(veiculo_blueprint, url_prefix='/veiculo')
app.register_blueprint(historico_blueprint, url_prefix='/historico')
app.register_blueprint(lembrete_blueprint, url_prefix='/lembrete')
