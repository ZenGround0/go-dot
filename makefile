all: go-dot

go-dot:
	go build

check:
	go vet ./...
	golint -set_exit_status -min_confidence 0.3 ./...

test:
	go test -v ./...
