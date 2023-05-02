FROM golang:1.20.3

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o watch-and-run .cmd/main.go

CMD ["./watch-and-run"]