package main

import (
	"fmt"
	"log"

	"github.com/353solutions/rides/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer conn.Close()

	log.Printf("info: connected to %s", addr)
	c := pb.NewRidesClient(conn)
	fmt.Println(c)
}
