FROM golang

EXPOSE 50111

ADD ./loadingServer /go/src/github.com/KitlerUA/GetMePhoto/loadingServer
ADD ./graber /go/src/github.com/KitlerUA/GetMePhoto/graber
RUN mkdir -p /home/kitler/Pictures
COPY ./images /home/kitler/Pictures
RUN go get github.com/Sirupsen/logrus
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
RUN go install github.com/KitlerUA/GetMePhoto/loadingServer

ENTRYPOINT ["/go/bin/loadingServer"]