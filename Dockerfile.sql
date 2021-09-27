# syntax=docker/dockerfile:1
FROM mysql
ENV MYSQL_ROOT_PASSWORD=123456 \
    MYSQL_DATABASE=chatroom

# docker-entrypoint-initdb.d will execute .sql files
ADD chatroom_chatroom_list.sql /docker-entrypoint-initdb.d 
ADD chatroom_userinfo.sql /docker-entrypoint-initdb.d

EXPOSE 3306
