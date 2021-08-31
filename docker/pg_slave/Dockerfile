FROM postgres:13-alpine3.14

RUN apk add --update su-exec
RUN apk add --update iputils

COPY ./docker/pg_slave/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["postgres"]

