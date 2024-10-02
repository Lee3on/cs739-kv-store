compose:
	docker compose up --build

compose_down:
	docker compose down

test:
	/bin/bash ./test.sh