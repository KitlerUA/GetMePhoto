FROM golang

EXPOSE 50112

ADD ./loadingServer /go/src/github.com/KitlerUA/GetMePhoto/loadingServer
ADD ./graber /go/src/github.com/KitlerUA/GetMePhoto/graber
ADD ./glide.yaml /go/src/github.com/KitlerUA/GetMePhoto

RUN mkdir -p /home/kitler/Pictures
COPY ./images /go/src/github.com/KitlerUA/GetMePhoto

RUN go get github.com/Masterminds/glide
WORKDIR /go/src/github.com/KitlerUA/GetMePhoto
RUN glide install

RUN go install github.com/KitlerUA/GetMePhoto/loadingServer

ENTRYPOINT ["/go/bin/loadingServer"]