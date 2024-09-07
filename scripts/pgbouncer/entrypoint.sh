#!/bin/sh
set -e

# Присвоим переменные окружения
DB_USER=${DB_USER}
DB_PASSWORD=${DB_PASSWORD}
DB_NAME=${DB_NAME}
DB_HOST=${DB_HOST}
DB_PORT=${DB_PORT} 
PGBOUNCER_PORT=${PGBOUNCER_PORT}

cat <<EOF > /etc/pgbouncer/userlist.txt
"${POSTGRES_USER}" "${POSTGRES_PASSWORD}"
EOF

cat <<EOF > /etc/pgbouncer/pgbouncer.ini
[databases]
${POSTGRES_DB} = host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD}

[pgbouncer]
listen_addr = *
listen_port = ${PGBOUNCER_PORT}
auth_type = md5
auth_file = /etc/pgbouncer/userlist.txt
pool_mode = session
max_client_conn = 100
default_pool_size = 20
EOF

echo "Configuration file /etc/pgbouncer/pgbouncer.ini created:"
cat /etc/pgbouncer/pgbouncer.ini

# Запуск PgBouncer
exec pgbouncer /etc/pgbouncer/pgbouncer.ini