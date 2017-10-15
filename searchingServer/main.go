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
)

const (
	defaultConnection = "localhost:50112"
)

func main() {
	port := flag.Int("p", 50111, "port to connect loading server")
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
	//fmt.Println(root)
	//err := filepath.Walk(root, visit)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Errorf("cannot read directory %v", root)
		return new(pg.Photo), err
	}
	var picturePath string = ""
	for _, f := range files {
		if !f.IsDir() {
			if strings.Contains(f.Name(), tag.Tag) {
				picturePath = filepath.Join(root, f.Name())
				break
			}
		}
	}

	fmt.Println(picturePath)

	conn, err := grpc.Dial(defaultConnection, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("cannot connect to seerver")
		return new(pg.Photo), err
	}

	client := pg.NewDownloadByIdClient(conn)
	id := &pg.Id{Url: picturePath}
	result, err := client.Download(context.Background(), id)
	if err != nil {
		fmt.Errorf("cannot download photo")
		return new(pg.Photo), err
	}

	return result, nil
}
