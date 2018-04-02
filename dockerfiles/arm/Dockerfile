FROM arm32v6/alpine
RUN apk --no-cache add ca-certificates
RUN mkdir www && mkdir files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
COPY asws /

WORKDIR /

CMD ["/asws"]