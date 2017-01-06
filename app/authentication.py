# -*- coding:utf-8 -*-
from functools import wraps
from flask import redirect, session,url_for

def auth_require():
    def wrapper(f):
        @wraps(f)
        def wrapped(*args, **kwargs):
            if not 'login' in session.keys():
                return redirect(url_for('login'))
            return f(*args, **kwargs)
        return wrapped
    return wrapper