FROM golang:1.10.2 AS aswsbuilder

RUN mkdir -p /go/src/github.com/txn2/asws
COPY . /go/src/github.com/txn2/asws

RUN go get ...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/asws ./src/github.com/txn2/asws

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
COPY --from=aswsbuilder /go/bin/asws /asws

WORKDIR /

ENTRYPOINT ["/asws"]