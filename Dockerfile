FROM golang:latest AS aswsbuilder

RUN mkdir -p /go/src/github.com/cjimti/asws
COPY . /go/src/github.com/cjimti/asws

RUN go get ...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/asws ./src/github.com/cjimti/asws

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=aswsbuilder /go/bin/asws /asws

WORKDIR /

ENTRYPOINT ["/asws"]