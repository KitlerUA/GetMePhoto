FROM golang

EXPOSE 50111

ADD ./searchingServer /go/src/github.com/KitlerUA/GetMePhoto/searchingServer
ADD ./graber /go/src/github.com/KitlerUA/GetMePhoto/graber
RUN mkdir -p /home/kitler/Pictures
COPY ./images /home/kitler/Pictures
RUN go get github.com/Sirupsen/logrus
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
RUN go install github.com/KitlerUA/GetMePhoto/searchingServer

ENTRYPOINT ["/go/bin/searchingServer"]