FROM arm32v6/alpine:3.6
RUN mkdir /www && mkdir /files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /

WORKDIR /

CMD ["/asws"]