FROM golang:1.19

#SET WORKING DIRECTORY
WORKDIR /src

#COPY CODE INTO WORKSPACE
COPY . .

#BUILD APP
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/servapi ./cmd/service/usermanagement
#REDUCE SIZE
RUN chmod +x -R /go

# API
COPY /_build/env/app/user-management-api-env.yml /src/config/config.yml

CMD ["/go/servapi"]



