# stage 1 build golang
FROM golang:alpine as builder

RUN apk update && apk add git && apk add make
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go build -o binary
# stage 2 reduce size
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/binary .

ENTRYPOINT [ "/app/binary" ]