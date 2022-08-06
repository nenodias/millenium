# -*- coding:utf-8 -*-
from functools import wraps
from flask import redirect, request, session, url_for, jsonify
from app import config
from app.utils import check_password, generate_hash


def auth_require():
    def wrapper(f):
        @wraps(f)
        def wrapped(*args, **kwargs):
            is_json = False
            if request.content_type and 'json' in request.content_type:
                is_json = True
            has_authorization = False
            if request.headers and 'Authorization' not in request.headers:
                has_authorization = True
            if is_json and has_authorization:
                authorization = request.headers.get('Authorization') or ''
                if not check_password(
                    config.DEFAULT_USERNAME,
                    config.DEFAULT_PASSWORD,
                    authorization
                ):
                    retorno = {"message": "must pass authorization header"}
                    return jsonify(retorno), 401
            elif not is_json and 'login' not in session.keys():
                return redirect(url_for('sistema.login'))
            return f(*args, **kwargs)
        return wrapped
    return wrapper
