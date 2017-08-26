import os
import logging

service = os.getenv('DATABASE_SERVICE_NAME', '').upper().replace('-', '_')

POSTGRESQL_DATABASE = os.getenv('DATABASE_NAME', 'millenium')
POSTGRESQL_DB_HOST = os.getenv(
    '{}_SERVICE_HOST'.format(POSTGRESQL_DATABASE),
    'localhost'
)
POSTGRESQL_DB_PORT = os.getenv(
    '{}_SERVICE_PORT'.format(POSTGRESQL_DATABASE),
    '5432'
)
POSTGRESQL_USER = os.getenv('DATABASE_USER', 'postgres')
POSTGRESQL_PASSWORD = os.getenv('DATABASE_PASSWORD', 'postgres')

DEFAULT_USERNAME = os.getenv('USER_DEFAULT', 'ADMIN')
DEFAULT_PASSWORD = os.getenv('PASS_DEFAULT', '123')

url = 'postgresql://{user}:{senha}'.format(
    user=POSTGRESQL_USER,
    senha=POSTGRESQL_PASSWORD
)

if POSTGRESQL_DB_HOST:
    url += '@{host}'.format(host=POSTGRESQL_DB_HOST)
if POSTGRESQL_DB_PORT:
    url += ':{port}'.format(port=POSTGRESQL_DB_PORT)
if POSTGRESQL_DATABASE:
    url += '/{database}'.format(database=POSTGRESQL_DATABASE)

url = os.getenv('DATABASE_URL', url)

logging.info(url)
SECRET_KEY = 'millenium'
SQLALCHEMY_DATABASE_URI = url
SQLALCHEMY_TRACK_MODIFICATIONS = True
