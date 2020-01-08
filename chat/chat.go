package chat

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/madjlzz/dummy-chat/api/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

const SeedLength = 42

type chatService struct {
	Subscribers map[string]*user
}

type user struct {
	username string
	stream   chat.ChatService_BroadcastServer
}

func NewChat() *chatService {
	return &chatService{
		Subscribers: make(map[string]*user),
	}
}

func (c *chatService) Broadcast(stream chat.ChatService_BroadcastServer) error {
	headers, _ := metadata.FromIncomingContext(stream.Context())
	token := headers["token"]
	user := c.Subscribers[token[0]]
	user.stream = stream

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("could not read request from client: %s\n", err))
		}

		for _, subscriber := range c.Subscribers {
			if subscriber.stream != stream && subscriber.stream != nil {
				err = subscriber.stream.Send(&chat.BroadcastResponse{
					Username: user.username,
					Content:  req.GetContent(),
				})
				if err != nil {
					return status.Errorf(
						codes.Internal,
						fmt.Sprintf("error while sending data to clients: %s\n", err))
				}
			}
		}
	}
}

func (c *chatService) Disconnect(ctx context.Context, req *chat.DisconnectRequest) (*chat.DisconnectResponse, error) {
	token := req.GetToken()

	if _, exist := c.Subscribers[token]; !exist {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("The user with token [%s] is already disconnected from the server.", token))
	}

	log.Printf("User (%s) has successfuly disconnected from the server.\n", c.Subscribers[token].username)
	delete(c.Subscribers, token)

	return &chat.DisconnectResponse{
		Disconnected: true,
	}, nil
}

func (c *chatService) Subscribe(ctx context.Context, req *chat.SubscribeRequest) (*chat.SubscribeResponse, error) {
	username := req.GetUsername()

	log.Printf("A new user (%s) is trying to join the chat.\n", username)

	for _, u := range c.Subscribers {
		if u.username == username {
			return nil, status.Errorf(
				codes.AlreadyExists,
				fmt.Sprintf("The user [%s] is already connected to the server.", username))
		}
	}

	token, err := generateToken()
	if err != nil {
		log.Printf("error during token generation: %s\n", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Something wrong happenned during user subscription."))
	}

	c.Subscribers[token] = &user{username: username}
	log.Printf("User (%s) has joined the chat!\n", username)

	return &chat.SubscribeResponse{
		Token: token,
	}, nil
}

func generateToken() (string, error) {
	t := make([]byte, SeedLength)
	_, err := rand.Read(t)
	return fmt.Sprintf("%x", t), err
}
