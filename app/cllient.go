package app

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	Peer string
	Name string
}

func NewClient(peer string, name string) *Client {
	return &Client{peer, name}
}

func (c *Client) SendMessage(content string) (*emptypb.Empty, error) {
	message := Message{From: c.Name, Content: []byte(content)}
	conn, err := grpc.Dial(c.Peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to %s: %s\n", c.Peer, err)
		return nil, err
	}
	defer conn.Close()

	cc := NewChatupApiClient(conn)
	_, err = cc.Send(context.Background(), &message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message to %s: %s\n", c.Peer, err)
		return nil, err
	}
	return nil, nil
}
