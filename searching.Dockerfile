FROM golang

EXPOSE 50111

ENV DWNLD=""

ADD ./searchingServer /go/src/github.com/KitlerUA/GetMePhoto/searchingServer
ADD ./graber /go/src/github.com/KitlerUA/GetMePhoto/graber
ADD ./glide.yaml /go/src/github.com/KitlerUA/GetMePhoto
RUN mkdir -p /home/kitler/Pictures
COPY ./images /home/kitler/Pictures

RUN go get github.com/Masterminds/glide
WORKDIR /go/src/github.com/KitlerUA/GetMePhoto
RUN glide install

RUN go install github.com/KitlerUA/GetMePhoto/searchingServer

ENTRYPOINT ["/go/bin/searchingServer"]