package main

import (
	"context"
	"fmt"

	presenceClient "github.com/SoroushBeigi/knowledge-game/adapter/presence"
	"github.com/SoroushBeigi/knowledge-game/dto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := presenceClient.New(conn)
	resp, err := client.GetPresence(
		context.Background(),
		dto.GetPresenceRequest{UserIDs: []uint{1, 2, 4}},
	)
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Items {
		fmt.Println("item", item.UserID, item.Timestamp)
	}
}
