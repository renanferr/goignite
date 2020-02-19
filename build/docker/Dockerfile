FROM registry.b2w.io/docker-run/base

COPY build/docker/app /usr/local/bin/app

ADD build/docker/run.sh /opt/run.sh

RUN chmod +x /usr/local/bin/app
RUN mkdir -p /opt/logs

WORKDIR /opt

CMD ["sh", "run.sh"]