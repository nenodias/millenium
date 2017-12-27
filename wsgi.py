"""WSGI."""
from app import create_app

application = create_app(False)

if __name__ == "__main__":
    application.run()
