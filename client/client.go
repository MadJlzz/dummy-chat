package main

import (
	"context"
	"flag"
	api "github.com/madjlzz/dummy-chat/api/chat"
	"google.golang.org/grpc"
	"log"
)

var (
	host     = flag.String("host", "127.0.0.1", "specify host for client to connect to")
	port     = flag.String("port", "50051", "specify port for client to connect to")
	username = flag.String("username", "anonymous", "client username used for server subscription")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*host+":"+*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to host [%s] on port [%s]: %s\n", *host, *port, err)
	}
	defer conn.Close()

	cli := api.NewChatClient(conn)
	_, err = cli.Subscribe(context.Background(), &api.SubscribeRequest{
		Username: *username,
	})
	if err != nil {
		log.Fatalf("error while trying to subscribe to server: %s\n", err)
	}

	log.Printf("Successfully connected to chat server [%s:%s] with username [%s].\n", *host, *port, *username)

	_, err = cli.Disconnect(context.Background(), &api.DisconnectRequest{
		Username: *username,
	})
	if err != nil {
		log.Fatalf("error while trying to disconnect from the server: %s\n", err)
	}

	log.Println("Successfully disconnected from chat server")
}
