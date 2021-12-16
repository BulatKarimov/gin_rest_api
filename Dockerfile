FROM golang:1.17.5-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /opt/code
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gin_rest_api cmd/api_server/main.go


FROM scratch

COPY --from=builder /opt/code/bin/gin_rest_api /usr/local/bin/gin_rest_api

ENTRYPOINT ["/usr/local/bin/gin_rest_api"]
