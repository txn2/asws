FROM alpine:3.18.3 as Builder

RUN mkdir -p /www && mkdir -p /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /bin/asws

RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/passwd

ENV PATH=/bin

WORKDIR /

ENV IP=0.0.0.0

USER nobody
ENTRYPOINT ["/bin/asws"]