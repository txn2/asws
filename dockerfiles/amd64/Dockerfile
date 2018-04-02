FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /

RUN mkdir www && mkdir files
COPY www/index.html www/index.html
COPY files/README.txt files/README.txt
CMD ["/asws"]