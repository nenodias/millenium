"""JOBS."""
import smtplib
from datetime import datetime
from email.message import EmailMessage


def create_email_job():
    """Method generate a method with no params to send mail."""
    from app import app, db
    from app.models import Lembrete

    def send_email():
        agora = datetime.now()
        lembretes = Lembrete.query.filter(
            Lembrete.data_notificacao <= agora
        ).all()
        print('Enviando emails')
        if lembretes:
            for lembrete in lembretes:
                email = EmailMessage()
                texto = lembrete.texto
                nome = ''
                veiculo = ''
                if lembrete.cliente is not None:
                    nome = lembrete.cliente.nome
                if lembrete.cliente is not None:
                    veiculo = lembrete.veiculo.descricao()

                mensagem = """
                {0}
                {1}
                {2}""".format(texto, nome, veiculo)
                email.set_content(mensagem)

                me = app.config['EMAIL_ME']
                you = app.config['EMAIL_YOU']
                password = app.config['EMAIL_ME_PASSWORD']
                smtp = app.config['EMAIL_SMTP']
                smtp_port = app.config['EMAIL_SMTP_PORT']

                email['Subject'] = 'Lembrete: {0}|{1}'.format(nome, veiculo)
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
