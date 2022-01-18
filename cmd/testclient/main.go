package main

import (
	"context"
	"log"

	"github.com/devnull-twitch/gameserver-manager/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewGameserverManagerClient(conn)

	response, err := client.GetGameserver(context.Background(), &proto.GetRequest{
		Zone: "overworld",
	})
	if err != nil {
		log.Fatalf("rpc error: %v", err)
	}

	log.Printf("%+v", response)
}
