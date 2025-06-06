FROM golang:1.24.3-alpine3.21 AS build-server

RUN apk update && \
    apk add git

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify

COPY . .
RUN go build -o ../bin/techlounge .

# -------------------------------------

FROM alpine:3.21

RUN addgroup -S techlounge && \
    adduser -S -D -H -h /home/techlounge -s /sbin/nologin -G techlounge techlounge && \
    mkdir -p /home/techlounge && \
    chown -R techlounge:techlounge /home/techlounge

WORKDIR /home/techlounge

COPY --from=build-server /go/bin/techlounge /usr/local/bin/techlounge

CMD [ "techlounge" ]
