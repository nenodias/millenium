# -*- coding: utf-8 -*-
from app import create_app

debug = False
app = create_app(debug=debug)

if __name__ == '__main__':    
    app.run(debug=debug, use_reloader=True, port=8080, host='0.0.0.0')
