package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	pg "github.com/KitlerUA/GetMePhoto/graber"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func main() {
	backend := flag.String("back", "172.18.0.3:50111", "server port")
	flag.Parse()

	ctx, cncl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cncl()
	conn, err := grpc.DialContext(ctx, *backend, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Cannot connect to serarching server: %v", err)
	}
	defer conn.Close()

	client := pg.NewGragPhotoClient(conn)
	tag := &pg.Tag{Tag: "Pira"}
	result, err := client.Get(context.Background(), tag)
	if err != nil {
		log.Fatalf("could not find image which contains %s: %v", tag.Tag, err)
	}
	if err = ioutil.WriteFile("temp.png", result.Image, 0666); err != nil {
		log.Fatalf("cannot write to file")
		return
	}
	fmt.Println("File created")
}
