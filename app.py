# -*- coding: utf-8 -*-
from app import app

if __name__ == '__main__':
    app.run(debug=False, use_reloader=True, port=8080, host='0.0.0.0')
