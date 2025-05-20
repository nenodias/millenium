FROM python:3.13

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8
ENV PORT=8000

WORKDIR /app

RUN useradd --create-home --home-dir /app --shell /bin/bash app
RUN chown -R app:app /app
RUN chmod -R 755 /app

ADD app  /app/app
ADD wsgi.py /app
ADD config.py /app
ADD requirements.txt /app
ADD setup.py /app

RUN apt-get update && apt-get install -y libpq-dev postgresql-client wget software-properties-common

RUN wget https://font.download/dl/font/arial.zip
RUN unzip arial.zip -d /usr/share/fonts/truetype/
RUN chmod 755 /usr/share/fonts/truetype/*
RUN rm -f arial.zip

RUN apt-get update && \
    apt-get install -y cabextract fontconfig && \
    fc-cache -f -v || true

RUN apt-get autoremove -y && \
    apt-get autoclean -y

USER app

RUN /usr/local/bin/python -m venv /app/venv
ENV PATH="/app/venv/bin:$PATH"
RUN /app/venv/bin/pip install --upgrade pip
RUN /app/venv/bin/pip install -r requirements.txt

EXPOSE 8000

ENTRYPOINT gunicorn wsgi:application --bind 0.0.0.0:$PORT