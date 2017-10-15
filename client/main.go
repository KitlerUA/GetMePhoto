package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	pg "github.com/KitlerUA/GetMePhoto/graber"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("back", "localhost:50111", "server port")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {

	}
	defer conn.Close()

	client := pg.NewGragPhotoClient(conn)
	tag := &pg.Tag{Tag: "pira"}
	result, err := client.Get(context.Background(), tag)
	if err != nil {
		log.Fatalf("could not say %s: %v", tag.Tag, err)
	}
	if err = ioutil.WriteFile("temp.png", result.Image, 0666); err != nil {
		fmt.Errorf("cannot write to file")
		return
	}
	fmt.Println("File created")
}
