FROM postgres:15.2

COPY /_build/script/pg/user-management-service/*.sql /docker-entrypoint-initdb.d/

# Set host time zone 
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone