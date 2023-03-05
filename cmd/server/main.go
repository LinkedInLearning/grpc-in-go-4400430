package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/353solutions/rides/pb"
)

func main() {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}

	srv := grpc.NewServer()
	var u Rides
	pb.RegisterRidesServer(srv, &u)
	reflection.Register(srv)

	log.Printf("info: server ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

func (r *Rides) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	// TODO: Validate req
	resp := pb.StartResponse{
		Id: req.Id,
	}

	// TODO: Work (insert to database ...)
	return &resp, nil
}

type Rides struct {
	pb.UnimplementedRidesServer
}
