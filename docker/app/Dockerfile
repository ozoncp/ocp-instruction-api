FROM golang:1.16.7-alpine3.14 AS build
RUN apk add --update make gcc libc-dev git
ENV PATH $PATH:/go/bin

WORKDIR /go/src
COPY ./api ./api
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./swagger ./swagger
COPY ./go.mod .
COPY ./go.sum .
COPY ./Makefile .

RUN go get -u github.com/pressly/goose/v3/cmd/goose
RUN make deps && make build


FROM alpine:3.14.1 AS bin
WORKDIR /root/
COPY --from=build /go/bin/goose ./
COPY --from=build /go/src/bin/ocp-instruction-api ./

COPY ./migrations ./migrations
COPY ./docker/app/migrate_and_run .
COPY ./config.yaml .

CMD ["./migrate_and_run"]
