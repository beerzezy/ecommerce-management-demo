FROM golang:1.19

#SET WORKING DIRECTORY
WORKDIR /src

#COPY CODE INTO WORKSPACE
COPY . .

#BUILD APP
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/servapi ./cmd/service/authentication
#REDUCE SIZE
RUN chmod +x -R /go

# API
COPY /_build/env/app/authentication-api-env.yml /src/config/config.yml

CMD ["/go/servapi"]



