"""JOBS."""
import threading
import smtplib
import pytz
from datetime import datetime
from email.mime.text import MIMEText

sao_paulo_tz = pytz.timezone('America/Sao_Paulo')


def create_email_job():
    """Method generate a method with no params to send mail."""
    from app import app, db
    from app.models import Lembrete

    lock = threading.Lock()

    def send_email():
        with lock:
            sp = datetime.now(tz=sao_paulo_tz)
            agora = datetime(
                year=sp.year,
                month=sp.month,
                day=sp.day,
                hour=sp.hour,
                minute=sp.minute
            )
            lembretes = Lembrete.query.filter(
                Lembrete.data_notificacao <= agora
            ).all()
            print('Enviando emails')
            if lembretes:
                for lembrete in lembretes:
                    texto = lembrete.texto
                    nome = ''
                    veiculo = ''
                    telefone = ''
                    celular = ''
                    tel_comercial = ''
                    e_mail = ''
                    if lembrete.cliente is not None:
                        nome = lembrete.cliente.nome
                        telefone = lembrete.cliente.telefone
                        celular = lembrete.cliente.celular
                        tel_comercial = lembrete.cliente.telefone_comercial
                        e_mail = lembrete.cliente.email
                    if lembrete.cliente is not None:
                        veiculo = lembrete.veiculo.descricao()

                    mensagem = """
                    Nome: {0}
                    Telefone: {1}
                    Celular: {2}
                    Telefone Comercial: {3}
                    E-mail: {4}
                    Veículo: {5}
                    Lembrete: {6}
                    """.format(
                        nome,
                        telefone,
                        celular,
                        tel_comercial,
                        e_mail,
                        veiculo,
                        texto
                    )
                    email = MIMEText(mensagem)

                    me = app.config['EMAIL_ME']
                    you = app.config['EMAIL_YOU']
                    password = app.config['EMAIL_ME_PASSWORD']
                    smtp = app.config['EMAIL_SMTP']
                    smtp_port = app.config['EMAIL_SMTP_PORT']

                    email['Subject'] = 'Lembrete: {0}|{1}'.format(
                        nome, veiculo
                    )
                    email['From'] = me
                    email['To'] = you

                    s = smtplib.SMTP(smtp, smtp_port)
                    s.login(me, password)
                    s.sendmail(me, [you], email.as_string())
                    s.quit()
                    # excluindo o lembrete
                    db.session.delete(lembrete)
                    db.session.commit()
    return send_email
