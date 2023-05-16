build:
	docker-compose build watch-and-run

run:
	docker-compose up watch-and-run

test:
	go test -v ./...

lint:
	golangci-lint run

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5433/postgres?sslmode=disable' up
