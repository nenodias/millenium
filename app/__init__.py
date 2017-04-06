# -*- coding: utf-8 -*-
from flask import (Flask, request, redirect, url_for, flash,
    jsonify,render_template, Blueprint, session)
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config['SECRET_KEY'] = 'millenium'
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://postgres:postgres@localhost:5432/millenium'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = True
db = SQLAlchemy(app)


from .authentication import auth_require
from app import controller