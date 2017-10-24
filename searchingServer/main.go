package main

import (
	"flag"
	"fmt"
	"net"

	"io/ioutil"

	"path/filepath"
	"strings"

	pg "github.com/KitlerUA/GetMePhoto/graber"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"os"
	"time"
)

const (
	defaultConnection = "downloading:50112"
	defaultPort       = 8080
)

func main() {
	port := flag.Int("p", defaultPort, "port to connect loading server")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("Cannot listen to %d: %v", *port, err)
	}

	serv := grpc.NewServer()
	pg.RegisterGragPhotoServer(serv, server{})

	err = serv.Serve(listener)
	if err != nil {
		logrus.Fatalf("Could not serve: %v", err)
	}
}

type server struct{}

func (server) Get(ctx context.Context, tag *pg.Tag) (*pg.Photo, error) {
	root := "/home/kitler/Pictures"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		logrus.Fatalf("cannot read directory %v", root)
		return new(pg.Photo), err
	}
	var picturePath string
	for _, f := range files {
		if !f.IsDir() {
			if strings.Contains(f.Name(), tag.Tag) {
				picturePath = filepath.Join(root, f.Name())
				break
			}
		}
	}

	var conn *grpc.ClientConn
	ctx, cncl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cncl()
	if address := os.Getenv("DWNLD"); address != "" {
		conn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		logrus.Printf("Address pass as arg: %v", address)
		if err != nil {
			logrus.Fatalf("cannot connect to server: %v", err)
		}
	} else {
		conn, err = grpc.DialContext(ctx, defaultConnection, grpc.WithInsecure(), grpc.WithBlock())
		logrus.Printf("Address use default address: %v", address)
		if err != nil {
			logrus.Fatalf("cannot connect to server: %v", err)
		}
	}

	client := pg.NewDownloadByIdClient(conn)
	id := &pg.Id{Url: picturePath}
	result, err := client.Download(context.Background(), id)
	if err != nil {
		logrus.Fatalf("cannot download photo: %v", err)
		return new(pg.Photo), err
	}
	return result, nil
}
