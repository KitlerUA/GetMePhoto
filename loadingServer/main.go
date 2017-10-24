package main

import (
	"flag"
	"net"

	"fmt"

	"io/ioutil"

	pg "github.com/KitlerUA/GetMePhoto/graber"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	defaultPort = 8080
)

func main() {
	port := flag.Int("p", defaultPort, "port to connect to loading server")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("Cannot listen to %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pg.RegisterDownloadByIdServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		logrus.Fatal("Could not serve")
	}
}

type server struct{}

func (server) Download(ctx context.Context, id *pg.Id) (*pg.Photo, error) {
	defer ctx.Done()
	photo, err := ioutil.ReadFile(id.Url)
	if err != nil {
		logrus.Fatalf("cannot read file %v", id.Url)
		return &pg.Photo{}, err
	}
	result := &pg.Photo{Image: photo}
	logrus.Infof(id.Url, "loaded")
	return result, nil
}
