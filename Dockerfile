FROM golang:1.6.3

MAINTAINER Nurul Arif Setiawan

# Change env here
# ENV GOBASE_COUCHBASEURI couchbase://192.168.99.101

RUN curl https://glide.sh/get | sh && \
    mkdir -p /go/src/github.com/arifsetiawan/gosimple 
WORKDIR /go/src/github.com/arifsetiawan/gosimple 

EXPOSE 4000

COPY . /go/src/github.com/arifsetiawan/gosimple
RUN glide install && go build && cp gosimple /usr/bin
CMD ["gosimple"]
