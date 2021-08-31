FROM postgres:13-alpine3.14

COPY ./docker/pg_master/setup-master.sh /docker-entrypoint-initdb.d/setup-master.sh
RUN chmod 0666 /docker-entrypoint-initdb.d/setup-master.sh

RUN mkdir /mnt/wal_arch
RUN chown -R postgres:postgres /mnt/wal_arch
