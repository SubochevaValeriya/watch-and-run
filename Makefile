build:
	docker-compose build watch-and-run

run:
	docker-compose up watch-and-run

test:
	go test -v ./...

lint:
	golangci-lint run
