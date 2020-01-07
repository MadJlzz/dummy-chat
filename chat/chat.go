package chat

import (
	"context"
	"fmt"
	"github.com/madjlzz/dummy-chat/api/chat"
	"github.com/madjlzz/dummy-chat/internal/structure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Chat struct {
	Subscribers structure.Set
}

func (c Chat) Disconnect(ctx context.Context, req *chat.DisconnectRequest) (*chat.DisconnectResponse, error) {
	username := req.GetUsername()

	log.Printf("User (%s) is trying disconnect.\n", username)

	if exist := c.Subscribers.Has(username); !exist {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("The user [%s] is already disconnected from the server.", username))
	}

	c.Subscribers.Remove(username)

	return &chat.DisconnectResponse{
		Disconnected: true,
	}, nil
}

func (c Chat) Subscribe(ctx context.Context, req *chat.SubscribeRequest) (*chat.SubscribeResponse, error) {
	username := req.GetUsername()

	log.Printf("A new user (%s) is trying to join the chat.\n", username)

	if exist := c.Subscribers.Has(username); exist {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("The user [%s] is already connected to the server.", username))
	}

	c.Subscribers.Add(username)

	return &chat.SubscribeResponse{
		Subscribed: true,
	}, nil
}
