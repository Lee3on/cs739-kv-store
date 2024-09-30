compose:
	docker-compose up --build

compose_down:
	docker-compose down

compose-go:
	docker-compose -f docker-compose_go.yml up --build
compose-go-down:
	docker-compose -f docker-compose_go.yml down

test:
	/bin/bash ./test.sh