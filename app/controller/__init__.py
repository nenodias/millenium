# -*- coding: utf-8 -*-
import pexpect
import sys, time
from flask import Response
from app import app, request, render_template, redirect, session, auth_require, url_for
from .falha_blueprint import falha_blueprint
from .peca_blueprint import peca_blueprint
from .servico_blueprint import servico_blueprint
from .tecnico_blueprint import tecnico_blueprint
from .montadora_blueprint import montadora_blueprint
from .modelo_blueprint import modelo_blueprint
from .cliente_blueprint import cliente_blueprint
from .veiculo_blueprint import veiculo_blueprint
from .historico_blueprint import historico_blueprint

@app.route('/login', methods=['POST', 'GET'])
def login():
    contexto = {}
    if request.method == 'POST':
        usuario = request.form.get('usuario')
        senha = request.form.get('senha')
        if usuario == 'ADMIN' and senha == '123':
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
@auth_require()
def index():
    return render_template('index.html'), 200

@app.route('/backup')
@auth_require()
def backup():
    uri = app.config['SQLALCHEMY_DATABASE_URI'].split('://')[1]
    parte = uri.split('@')
    usuario, senha = parte[0].split(':')
    parte = parte[1].split(':')
    host = parte[0]
    porta, database = parte[1].split('/')
    # executando o pg_dump
    call = 'pg_dump -d '+database+' -p '+porta+' -U '+usuario+' -h '+host+' -W'
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
