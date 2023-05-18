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

mock:
	mockgen --destination mocks/event_mock.go --package=mocks -source internal/repository/event.go
	mockgen --destination mocks/launch_mock.go --package=mocks -source internal/repository/launch.go
