from datetime import datetime
import time

DATETIME_FORMAT = '%d/%m/%Y %H:%M:%S'

def from_datetime_to_str(value):
	if value:
		return value.strftime(DATETIME_FORMAT)
	return ''

def from_str_to_datetime_or_none(value):
	try:
		return datetime.strptime(value, DATETIME_FORMAT)
	except:
		return None


def to_int_or_none(value):
    try:
        return int(value)
    except:
        pass
    return None