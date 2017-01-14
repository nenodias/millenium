from datetime import datetime
import time

HTML_DATETIME_FORMAT = '%Y-%m-%dT%H:%M'
DATETIME_FORMAT = '%d/%m/%Y %H:%M:%S'
HTML_DATE_FORMAT = '%Y-%m-%d'
DATE_FORMAT = '%d/%m/%Y'
FINAL_MOMENT_OF_DAY = '23:59:59'

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

def final_date_day(value):
	date = value.strftime(DATE_FORMAT)
	date += ' '+ FINAL_MOMENT_OF_DAY
	return from_str_to_datetime_or_none(date)

def to_int_or_none(value):
    try:
        return int(value)
    except:
        pass
    return None