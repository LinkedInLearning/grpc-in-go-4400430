package rides

/* Get the Go gRPC tools by running

	$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

You need the tools to be in your path:

	$ export PATH="$PATH:$(go env GOPATH)/bin"

See more at https://grpc.io/docs/languages/go/quickstart/
*/

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative ./rides.proto
//go:generate protoc --grpc-gateway_out pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./rides.proto
