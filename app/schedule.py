"""Agenda."""
import sys
import threading
import time
from .email import send_email
from .utils import final_date_day, get_now

CONTACT_MAIL = 'horacio.dias@yahoo.com'

def enviar(lembrete):
    message = lembrete.texto
    subject = "Lembrete: %s" % (lembrete.id)
    send_email(CONTACT_MAIL, subject, message)


def start():
    from .models import db, Lembrete

    def enviar_lembrete():
        _data = get_now()
        _end = final_date_day(_data)
        result = Lembrete.query.filter(
            Lembrete.data_notificacao.between(_data, _end)).all()
        if result:
            for lembrete in result:
                print('Enviando lembrete: %s' % (lembrete.id))
                enviar(lembrete)
                db.session.delete(lembrete)
                db.session.commit()
        time.sleep(1000 * 60)

    try:
        t = threading.Thread(target=enviar_lembrete, args=())
        t.start()
    except Exception as ex:
        sys.exit(0)
