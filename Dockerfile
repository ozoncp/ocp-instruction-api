FROM golang:1.16.7-alpine3.14 AS build
RUN apk add --update make gcc libc-dev git
ENV PATH $PATH:/go/bin
WORKDIR /go/src
COPY . .
RUN go get -u github.com/pressly/goose/v3/cmd/goose
RUN make deps && make build


FROM alpine:3.14.1 AS bin
WORKDIR /root/
COPY --from=build /go/bin/goose ./
COPY --from=build /go/src/migrations ./
COPY --from=build /go/src/bin/ocp-instruction-api ./
COPY --from=build /go/src/migrate_and_run ./

CMD ["./migrate_and_run"]
