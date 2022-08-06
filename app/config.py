"""Config."""
import os
import binascii

DEFAULT_USERNAME = os.getenv('USER_DEFAULT', 'ADMIN')
DEFAULT_PASSWORD = os.getenv('PASS_DEFAULT', '123')

POSTGRES_DEFAULT_USER = 'postgres'
POSTGRES_USER = os.getenv('POSTGRES_USER', 'dbuser')
POSTGRES_PASSWORD = os.getenv('POSTGRES_PASSWORD', 'dbpassword')
POSTGRES_DB = os.getenv('POSTGRES_DB', 'dbsample')

PORT = os.getenv('POSTGRESQL_SERVICE_PORT', '5432')
HOST = os.getenv('POSTGRESQL_SERVICE_HOST', 'postgres')

DB = '{user}:{passw}@{host}:{port}/{database}'.format(
    user=POSTGRES_USER,
    passw=POSTGRES_PASSWORD,
    host=HOST,
    port=PORT,
    database=POSTGRES_DB
)
DATABASE_URI = 'postgresql://' + DB

SECRET_KEY = os.getenv('SERVER_SECRET', binascii.hexlify(os.urandom(24)))
SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', DATABASE_URI)
SQLALCHEMY_TRACK_MODIFICATIONS = True

# email config
EMAIL_ME = os.getenv('EMAIL_ME', 'from@teste.com.br')
EMAIL_ME_PASSWORD = os.getenv('EMAIL_ME_PASSWORD', 'senha')
EMAIL_YOU = os.getenv('EMAIL_YOU', 'you@teste.com.br')
EMAIL_SMTP = os.getenv('EMAIL_SMTP', 'smtps.teste.com.br')
EMAIL_SMTP_PORT = os.getenv('EMAIL_SMTP_PORT', 587)
SEND_FILE_MAX_AGE_DEFAULT = 0