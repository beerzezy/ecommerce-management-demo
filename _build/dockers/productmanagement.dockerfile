FROM golang:1.19

#SET WORKING DIRECTORY
WORKDIR /src

#COPY CODE INTO WORKSPACE
COPY . .

#BUILD APP
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/servapi ./cmd/servapi

# API
COPY /_build/env/authentication-app-env.yml /src/config/config.yml
COPY /_build/env/order-management-app-env.yml /src/config/config.yml
COPY /_build/env/product-management-app-env.yml /src/config/config.yml
COPY /_build/env/user-management-app-env.yml /src/config/config.yml



#REDUCE SIZE
RUN chmod +x -R /go

CMD ["/go/servapi"]

FROM postgres:15.2

COPY /_build/script/pg/*.sql /docker-entrypoint-initdb.d/

# Set host time zone 
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
