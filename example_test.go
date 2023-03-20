package backoffice

import (
	"fmt"
	"time"

	"github.com/353solutions/backoffice/pb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExampleInvoice() {
	// TODO: Change to protobuf invoice
	time := time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)
	inv := pb.Invoice{
		Id:       "2023-0123",
		Time:     timestamppb.New(time),
		Customer: "Wile E. Coyote",
		Items: []*pb.LineItem{
			{Sku: "hammer-20", Amount: 1, Price: 249},
			{Sku: "nail-9", Amount: 100, Price: 1},
			{Sku: "glue-5", Amount: 1, Price: 799},
		},
	}
	fmt.Printf("%v\n", &inv) // Make compiler happy
	// TODO: Encode to []byte using protobuf
	data, err := proto.Marshal(&inv)
	if err == nil {
		fmt.Println("size:", len(data))
	} else {
		fmt.Println("ERROR:", err)
	}

	// Output:
	// id:"2023-0123" time:{seconds:1673099100} customer:"Wile E. Coyote" items:{sku:"hammer-20" amount:1 price:249} items:{sku:"nail-9" amount:100 price:1} items:{sku:"glue-5" amount:1 price:799}
	// size: 82
}
