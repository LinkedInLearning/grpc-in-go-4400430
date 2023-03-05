package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/353solutions/rides/pb"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"localhost:9292",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()
	err = pb.RegisterRidesHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}

	addr := ":8080"
	log.Printf("gateway server starting on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
