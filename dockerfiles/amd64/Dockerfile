FROM alpine:3.7
RUN mkdir /www && mkdir /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /

WORKDIR /

CMD ["/asws"]