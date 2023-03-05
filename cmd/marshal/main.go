package main

import (
	"fmt"
	"log"

	"github.com/353solutions/rides/pb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	req := pb.StartRequest{
		Id:       "47a74960d6204a52b1bece53221eb458",
		DriverId: "007",
		Location: &pb.Location{
			Lat: 51.4871871,
			Lng: -0.1266743,
		},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}
	fmt.Println(&req)

	data, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var req2 pb.StartRequest
	if err := proto.Unmarshal(data, &req2); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(&req2)
}
