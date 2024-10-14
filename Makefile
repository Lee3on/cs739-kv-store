compose:
	docker compose up --build

compose_down:
	docker compose down

test:
	/bin/bash ./test.sh

gen_go:
	protoc \
    	--go_out=./server/proto/kv739 \
    	--go_opt=module=cs739-kv-store/proto/kv739 \
    	--go-grpc_out=./server/proto/kv739 --go-grpc_opt=module=cs739-kv-store/proto/kv739 \
    	proto/kv739.proto
	cp ./server/proto/kv739/* ./load_balancer/proto/kv739/