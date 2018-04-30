"""App."""
import logging
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_apscheduler import APScheduler
from app import jobs

db = SQLAlchemy()

def create_app(debug=False):
    """Cria a aplicacao."""
    app = Flask(__name__)
    app.debug = debug
    app.config.from_pyfile('config.py')
    from .authentication import auth_require
    from app import controller
    logging.info(auth_require)
    logging.info(controller)
    controller.init_app(app)
    db.app = app
    db.init_app(app)

    send_mail = jobs.create_email_job(app, db)

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
    scheduler = APScheduler()
    scheduler.init_app(app)
    scheduler.start()

    return app
