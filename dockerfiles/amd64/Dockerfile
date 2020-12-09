FROM alpine:3.12.1 as Builder

RUN mkdir -p /www && mkdir -p /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /bin/asws

RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

WORKDIR /

FROM scratch

ENV PATH=/bin

WORKDIR /

ENV IP=0.0.0.0

COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder /www/index.html /www/index.html
COPY --from=builder /files/README.txt /files/README.txt
COPY --from=builder /bin/asws /bin/asws

USER nobody
ENTRYPOINT ["/bin/asws"]