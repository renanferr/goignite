FROM golang:1.13-buster

COPY build/docker/app /usr/local/bin/app

ADD build/docker/run.sh /opt/run.sh

RUN chmod +x /usr/local/bin/app
RUN mkdir -p /opt/logs

WORKDIR /opt

CMD ["sh", "run.sh"]