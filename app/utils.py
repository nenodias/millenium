def to_int_or_none(value):
    try:
        return int(value)
    except:
        pass
    return None