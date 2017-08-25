# -*- coding: utf-8 -*-
from app import app, scheduler

if __name__ == '__main__':
    scheduler.start()
    app.run(debug=False, use_reloader=True, port=8080, host='0.0.0.0')
