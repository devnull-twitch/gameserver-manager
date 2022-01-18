package main

import (
	"fmt"
	"log"
	"net"

	internal_grpc "github.com/devnull-twitch/gameserver-manager/lib/grpc"
	"github.com/devnull-twitch/gameserver-manager/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	rpcServerImplementation := internal_grpc.GetServer()
	proto.RegisterGameserverManagerServer(s, rpcServerImplementation)

	defer func() {
		rpcServerImplementation.StopServer()
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
