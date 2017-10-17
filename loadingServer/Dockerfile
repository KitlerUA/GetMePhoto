FROM golang

ADD . /go/src/github.com/KitlerUA/GetMePhoto

RUN go get github.com/Sirupsen/logrus

RUN go get golang.org/x/net/context

RUN go get google.golang.org/grpc

RUN go install github.com/KitlerUA/GetMePhoto/loadingServer

ENTRYPOINT ["/go/bin/loadingServer"]

EXPOSE 50112
