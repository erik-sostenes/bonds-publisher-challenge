server:
	bash ./scripts/build.bash

keys:
	bash ./scripts/keys.bash

test:
	go test -v ./...\

build:
	docker-compose up --build

migrate:
	bash ./scripts/migrate.bash

run: server
	./build