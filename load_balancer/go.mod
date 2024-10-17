module load_balancer

go 1.18

require (
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.34.2
)

replace cs739-kv-store => ../server

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
)