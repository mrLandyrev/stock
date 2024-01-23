FROM postgres:15
COPY ./dump-example.sql /docker-entrypoint-initdb.d/