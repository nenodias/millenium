FROM python:3.9

ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
ENV PORT 8000

WORKDIR /app

ADD app  /app/app
ADD wsgi.py /app
ADD config.py /app
ADD requirements.txt /app
ADD setup.py /app

RUN /usr/local/bin/pip install --upgrade pip
RUN /usr/local/bin/pip install -r requirements.txt

EXPOSE 8000

ENTRYPOINT gunicorn wsgi:application --bind 0.0.0.0:$PORT