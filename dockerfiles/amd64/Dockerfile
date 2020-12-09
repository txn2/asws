FROM alpine:3.12.1
RUN mkdir -p /www && mkdir -p /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /

WORKDIR /
#RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/passwd
USER nobody

CMD ["/asws"]