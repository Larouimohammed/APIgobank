build:
@go build -o bin/Gobank
run : build
@.bin/Gobank
test:
@go test -v ./...
