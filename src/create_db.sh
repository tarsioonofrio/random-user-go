source .env
export $(cut -d= -f1 .env)

psql $PGCONN -c "CREATE DATABASE randomuser"
