FROM golang:1.17.5-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/gin_rest_api cmd/api_server/main.go
ENTRYPOINT ["bin/gin_rest_api"]