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

run_bash:
	docker exec -it kv739 /bin/bash
build_image: clean
	docker build -t kv739 .

run_image: build_image
	docker run -it --name kv739 -p 50051:50051 kv739
exec_test:
	docker exec -it kv739 ./kv739_test
put:
	docker exec -it kv739 ./kv739_test --put
clean:
	docker rm kv739 -f
	docker rmi kv739:latest -f
	docker image prune -f