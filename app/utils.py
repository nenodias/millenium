import hashlib
import time
from datetime import datetime, date
from werkzeug.security import generate_password_hash, \
     check_password_hash

HTML_DATETIME_FORMAT = '%Y-%m-%dT%H:%M'
DATETIME_FORMAT = '%d/%m/%Y %H:%M:%S'
HTML_DATE_FORMAT = '%Y-%m-%d'
DATE_FORMAT = '%d/%m/%Y'
FINAL_MOMENT_OF_DAY = '23:59:59'
START_MOMENT_OF_DAY = '00:00:00'


def from_datetime_to_str(value):
    if value:
        return value.strftime(HTML_DATETIME_FORMAT)
    return ''


def from_str_to_datetime_or_none(value):
    try:
        return datetime.strptime(value, HTML_DATETIME_FORMAT)
    except:
        return None


def from_str_to_date_or_none(value):
    try:
        return datetime.strptime(value, HTML_DATE_FORMAT)
    except:
        return None


def format_date(value):
    try:
        return datetime.strftime(value, DATE_FORMAT)
    except:
        return ''


def final_date_day(value):
    date = value.strftime(DATE_FORMAT)
    date += ' ' + FINAL_MOMENT_OF_DAY
    return datetime.strptime(date, DATETIME_FORMAT)


def get_now():
    value = datetime.now()
    date = value.strftime(DATE_FORMAT)
    date += ' ' + START_MOMENT_OF_DAY
    return datetime.strptime(date, DATETIME_FORMAT)


def to_int_or_none(value):
    try:
        return int(value)
    except:
        pass
    return None


def to_float_or_zero(value):
    try:
        return float(value)
    except:
        pass
    return 0


def generate_hash(usuario, senha):
    str_hash = format_date(date.today()) + usuario+senha
    h = hashlib.md5(bytes(str_hash, 'utf-8'))
    return h.hexdigest()


def check_password(usuario, senha, password_hash):
    hashed = generate_hash(usuario, senha)
    return hashed == password_hash
