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
	"log"
)

func main() {
	port := flag.Int("p", 50112, "port to connect to loading server")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("Cannot listen to %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pg.RegisterDownloadByIdServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		fmt.Errorf("could not serve")
	}
}

type server struct{}

func (server) Download(ctx context.Context, id *pg.Id) (*pg.Photo, error) {
	log.Println("In Download func")
	defer ctx.Done()
	photo, err := ioutil.ReadFile(id.Url)
	if err != nil {
		fmt.Errorf("cannot read file %v", id.Url)
		return &pg.Photo{}, err
	}
	result := &pg.Photo{Image: photo}
	log.Println(id.Url, "loaded")
	return result, nil
}
