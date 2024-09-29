gen:
	protoc \
	--go_out=./proto/kv739 \
	--go_opt=module=cs739-kv-store/proto/kv739 \
	--go-grpc_out=./proto/kv739 --go-grpc_opt=module=cs739-kv-store/proto/kv739 \
	proto/kv739/kv739.proto

run:
	go run server.go handler.go -port 6666
