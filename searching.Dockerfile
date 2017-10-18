FROM golang

EXPOSE 50111

ADD ./loadingServer /go/src/github.com/KitlerUA/GetMePhoto/loadingServer
ADD ./graber /go/src/github.com/KitlerUA/GetMePhoto/graber
RUN go get github.com/Sirupsen/logrus
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
RUN go install github.com/KitlerUA/GetMePhoto/loadingServer

ENTRYPOINT ["/go/bin/loadingServer"]