package app

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	UnimplementedChatupApiServer
}

func (s *service) Send(ctx context.Context, message *Message) (*emptypb.Empty, error) {
	fmt.Fprintf(os.Stdout, "[%s] %s\n", message.From, message.Content)
	return nil, nil
}

type Server struct {
	Uri string
}

func NewServer(uri string) *Server {
	return &Server{uri}
}

func (s *Server) Listen() error {
	lis, err := net.Listen("tcp", s.Uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting tcp server: %s\n", err)
		return err
	}

	server := grpc.NewServer()
	RegisterChatupApiServer(server, &service{})

	return server.Serve(lis)
}
