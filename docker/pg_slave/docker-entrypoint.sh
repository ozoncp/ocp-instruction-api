#!/bin/bash

if [ ! -s "$PGDATA/PG_VERSION" ]; then
  echo "*:*:*:$PG_REP_USER:$PG_REP_PASSWORD" > ~/.pgpass
  chmod 0600 ~/.pgpass

  until ping -c 1 -W 1 pg_master
  do
    echo "Waiting for master to ping..."
    sleep 1s
  done

  until pg_basebackup -h pg_master -D ${PGDATA} -U ${PG_REP_USER} -vP -W
  do
    echo "Waiting for master to connect..."
    sleep 1s
  done

  cat >> ${PGDATA}/postgresql.conf <<EOF
    primary_conninfo = 'host=pg_master port=5432 user=$PG_REP_USER password=$PG_REP_PASSWORD'
    restore_command = 'cp /mnt/wal_arch/%f %p'
    archive_cleanup_command = 'pg_archivecleanup /mnt/wal_arch %r'
EOF

  echo "" > ${PGDATA}/standby.signal

  chown postgres. ${PGDATA} -R
  chmod 700 ${PGDATA} -R
fi

exec su-exec postgres "$@"
