package chat

import (
	"context"
	"github.com/madjlzz/dummy-chat/api/chat"
)

type Chat struct {
}

func (c Chat) Subscribe(ctx context.Context, req *chat.SubscribeRequest) (*chat.SubscribeResponse, error) {
	panic("implement me")
}
