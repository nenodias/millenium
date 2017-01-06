# -*- coding: utf-8 -*-
from app import app, request, render_template, redirect, session, auth_require, url_for
from .falha_blueprint import falha_blueprint

@app.route('/login', methods=['POST', 'GET'])
def login():
    contexto = {}
    if request.method == 'POST':
        usuario = request.form.get('usuario')
        senha = request.form.get('senha')
        if usuario == 'admin' and senha == '123':
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

app.register_blueprint(falha_blueprint, url_prefix='/falha')