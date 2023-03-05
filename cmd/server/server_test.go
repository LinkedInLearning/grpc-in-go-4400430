package main

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/353solutions/rides/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestEnd(t *testing.T) {
	req := pb.EndRequest{
		Id:       "end",
		Time:     timestamppb.Now(),
		Distance: 3.14,
	}

	var srv Rides
	resp, err := srv.End(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Id != req.Id {
		t.Fatalf("bad response id: got %#v, expected %#v", resp.Id, req.Id)
	}
}

func TestEndE2E(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	srv := createServer()
	go srv.Serve(lis)

	port := lis.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("localhost:%d", port)
	creds := insecure.NewCredentials()
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewRidesClient(conn)

	req := pb.EndRequest{
		Id:       "end",
		Time:     timestamppb.Now(),
		Distance: 3.14,
	}

	resp, err := c.End(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Id != req.Id {
		t.Fatalf("bad response id: got %#v, expected %#v", resp.Id, req.Id)
	}
}
