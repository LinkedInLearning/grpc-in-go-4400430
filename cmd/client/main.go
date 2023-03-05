package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/353solutions/rides/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer conn.Close()

	log.Printf("info: connected to %s", addr)
	c := pb.NewRidesClient(conn)
	fmt.Println(c)
}
