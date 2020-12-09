FROM alpine:3.12.1 as Builder

RUN mkdir -p /www && mkdir -p /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /bin/asws

WORKDIR /

FROM scratch

ENV PATH=/bin

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder /www/index.html /www/index.html
COPY --from=builder /files/README.txt /files/README.txt
COPY --from=builder /bin/asws /bin/asws

USER nobody
ENTRYPOINT ["/bin/asws"]