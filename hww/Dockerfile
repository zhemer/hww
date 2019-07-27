FROM golang:alpine
MAINTAINER Sergey Zhemerdeev zhemers@gmail.com https://github.com/zhemer/hww/tree/master/hww
RUN apk add curl
COPY hww.go /
WORKDIR /
RUN go build -o /hww
#EXPOSE 1234
CMD /hww -request-log 1
HEALTHCHECK --interval=10s CMD curl -f http://localhost:8080/healthz || exit 1
