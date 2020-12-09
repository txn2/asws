FROM alpine:3.12.1
RUN mkdir -p /www && mkdir -p /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /bin/asws

WORKDIR /
USER nobody

ENTRYPOINT ["/bin/asws"]