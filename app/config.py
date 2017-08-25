import os

service = os.getenv('DATABASE_SERVICE_NAME', '').upper().replace('-', '_')

POSTGRESQL_DATABASE = os.getenv('DATABASE_NAME')
POSTGRESQL_DB_HOST = os.getenv('{}_SERVICE_HOST'.format(POSTGRESQL_DATABASE))
POSTGRESQL_DB_PORT = os.getenv('{}_SERVICE_PORT'.format(POSTGRESQL_DATABASE))
POSTGRESQL_USER = os.getenv('DATABASE_USER')
POSTGRESQL_PASSWORD = os.getenv('DATABASE_PASSWORD')


DEFAULT_USERNAME = os.getenv('USER_DEFAULT', 'ADMIN')
DEFAULT_PASSWORD = os.getenv('PASS_DEFAULT','123')
'''
POSTGRESQL_DATABASE = os.getenv('DATABASE_NAME','millenium')
POSTGRESQL_USER = os.getenv('POSTGRESQL_USER', 'postgres')
POSTGRESQL_PASSWORD = os.getenv('POSTGRESQL_PASSWORD', 'postgres')
POSTGRESQL_DB_HOST = os.getenv('OPENSHIFT_POSTGRESQL_DB_HOST', 'localhost')
POSTGRESQL_DB_PORT = os.getenv('OPENSHIFT_POSTGRESQL_DB_PORT','5432')
'''
url = 'postgresql://{user}:{senha}@{host}'.format(
    user=POSTGRESQL_USER,
    senha=POSTGRESQL_PASSWORD,
    host=POSTGRESQL_DB_HOST
)
if POSTGRESQL_DB_PORT:
    url += ':{port}'.format(port=POSTGRESQL_DB_PORT)
if POSTGRESQL_DATABASE:
    url += '/{database}'.format(database=POSTGRESQL_DATABASE)

SECRET_KEY = 'millenium'
SQLALCHEMY_DATABASE_URI = url
SQLALCHEMY_TRACK_MODIFICATIONS = True
