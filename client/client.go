package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	api "github.com/madjlzz/dummy-chat/api/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"os/signal"
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

	cli := api.NewChatServiceClient(conn)
	resp, err := cli.Subscribe(context.Background(), &api.SubscribeRequest{
		Username: *username,
	})
	if err != nil {
		log.Fatalf("error while trying to subscribe to server: %s\n", err)
	}

	token := resp.GetToken()
	log.Printf("Successfully connected to chat server [%s:%s] with username [%s].\n", *host, *port, *username)
	log.Printf("ID token received after subscription [%s]\n", token)

	m := metadata.Pairs("token", token)
	ctx := metadata.NewOutgoingContext(context.Background(), m)

	stream, err := cli.Broadcast(ctx)
	if err != nil {
		log.Fatalf("error while streaming incoming and outgoing messages: %s\n", err)
	}

	fmt.Println()

	go captureInput(token, stream)
	go receiveMessages(stream)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Printf("Trying to disconnect the client [%s]\n", *username)
	_, err = cli.Disconnect(context.Background(), &api.DisconnectRequest{
		Token: token,
	})
	if err != nil {
		log.Fatalf("error while trying to disconnect from the server: %s\n", err)
	}
	log.Println("Successfully disconnected from chat server")
}

// Basically a user should be able to input text from stdin
// and send them to the server...
func captureInput(token string, stream api.ChatService_BroadcastClient) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You're saying: ")
		input, err := r.ReadString('\n')
		if err != nil {
			log.Printf("error while trying to read stdin: %s\n", err)
		}
		err = stream.Send(&api.BroadcastRequest{
			Token:   token,
			Content: input,
		})
		if err != nil {
			log.Printf("error while sending message to server: %s\n", err)
		}
	}
}

// Any time the server sends a message, log it
// to the client with it's username.
func receiveMessages(stream api.ChatService_BroadcastClient) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			log.Printf("error while receiving message from server: %s\n", err)
		}
		fmt.Println()
		fmt.Printf("%s says: %s\n", resp.GetUsername(), resp.GetContent())
		fmt.Print("You're saying: ")
	}
}
