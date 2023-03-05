package main

import (
	"fmt"

	"github.com/353solutions/rides/pb"

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
	fmt.Println(req)
}
