FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD ${MYSQL_ROOT_PASSWORD}
ENV MYSQL_DATABASE ${MYSQL_DATABASE}

# RUN mysql
# RUN create database ultividros_db;

EXPOSE ${MYSQL_PORT}
