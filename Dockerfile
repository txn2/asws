ARG app=asws
ARG project=github.com/txn2/asws
ARG buildsrc=./cmd/asws.go

FROM golang:1.15.2-alpine3.12 AS builder

ARG app
ARG project
ARG buildsrc
ARG version

ENV PROJECT=${project} \
    APP=${app} \
    BUILDSRC=${buildsrc} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /go/src/ \
 && mkdir -p /go/bin \
 && mkdir -p /go/pkg

ENV PATH=/go/bin:$PATH

RUN mkdir -p /go/src/$PROJECT/
ADD . /go/src/$PROJECT/

WORKDIR /go/src/$PROJECT/

RUN go build -ldflags "-X main.Version=${version} -extldflags \"-static\"" -o /go/bin/app $BUILDSRC
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

RUN mkdir -p /www && mkdir -p /files
COPY www/index.html /www/index.html
COPY files/README.txt /files/README.txt

FROM scratch

ENV PATH=/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder /www/index.html /www/index.html
COPY --from=builder /files/README.txt /files/README.txt
COPY --from=builder /go/bin/app /bin/asws

WORKDIR /

USER nobody
ENTRYPOINT ["/bin/asws"]