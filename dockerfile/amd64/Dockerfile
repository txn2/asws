FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /

RUN mkdir www && mkdir files
COPY index.html www/index.html
COPY README.txt files/README.txt
CMD ["/asws"]