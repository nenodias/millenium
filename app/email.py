import smtplib
from email.mime.text import MIMEText
try:
    from email.message import EmailMessage
except:
    pass


def send_email(you, subject, content):
    try:
        msg = EmailMessage()
        msg.set_content(content)
    except:
        msg = MIMEText(content)

    msg['Subject'] = subject
    msg['From'] = ME
    msg['To'] = you

    s = smtplib.SMTP(MAIL_SERVER, 587)
    s.login(ME, PASSWD)
    s.sendmail(ME, [you], msg.as_string())
    s.quit()
