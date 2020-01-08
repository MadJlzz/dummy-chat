package main

import (
	"flag"
	api "github.com/madjlzz/dummy-chat/api/chat"
	"github.com/madjlzz/dummy-chat/chat"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	port = flag.String("port", "50051", "specify port for server to listen on")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	lis, err := net.Listen("tcp", "127.0.0.1:"+*port)
	if err != nil {
		log.Fatalf("error occured while trying to listen on port [%s]: %s\n", *port, err)
	}

	s := grpc.NewServer()
	api.RegisterChatServiceServer(s, chat.NewChat())

	go func() {
		log.Printf("Starting Server on port [%s]...\n", *port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("error occured while trying to serve on port [%s]: %s\n", *port, err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Stopping the Server gracefully...")
	s.Stop()

	log.Println("Closing the listener gracefully...")
	_ = lis.Close()
}
