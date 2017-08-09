# -*- coding: utf-8 -*-
from flask import (Flask, request, redirect, url_for, flash,
    jsonify, render_template, Blueprint, session)
from flask_sqlalchemy import SQLAlchemy
from flask_apscheduler import APScheduler
from app import jobs

app = Flask(__name__)
app.config['SECRET_KEY'] = 'millenium'
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://postgres:postgres@localhost:5432/millenium'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = True

db = SQLAlchemy(app)

from .authentication import auth_require
from app import controller

send_mail = jobs.create_email_job()

app.config['JOBS'] = [
    {
        'id': 'send_mail',
        'func': send_mail,
        'trigger': 'interval',
        'seconds': 60
    }
]
app.config['SCHEDULER_API_ENABLED'] = True
app.config['SCHEDULER_EXECUTORS'] = {
    'default': {'type': 'threadpool', 'max_workers': 1}
}
# email config
app.config['EMAIL_ME'] = 'from@teste.com.br'
app.config['EMAIL_ME_PASSWORD'] = 'senha'
app.config['EMAIL_YOU'] = 'you@teste.com.br'
app.config['EMAIL_SMTP'] = 'smtps.teste.com.br'
app.config['EMAIL_SMTP_PORT'] = 587

scheduler = APScheduler()
scheduler.init_app(app)
