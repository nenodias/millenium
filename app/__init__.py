# -*- coding: utf-8 -*-
from flask import (
    Flask, request, redirect, url_for, flash,
    jsonify, render_template, Blueprint, session
)
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config.from_pyfile('config.py')
db = SQLAlchemy(app)


from .authentication import auth_require
from app import controller
