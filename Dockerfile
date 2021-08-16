FROM golang:1.16.7-alpine3.14 AS build
RUN apk add --update make
ENV PATH $PATH:/go/bin
WORKDIR /go/src
COPY . .
RUN make deps && make build


FROM alpine:3.14.1 AS bin
WORKDIR /root/
COPY --from=build /go/src/bin/ocp-instruction-api ./
CMD ["./ocp-instruction-api"]
