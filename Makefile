build:
	mkdir -p bin
	go build -o bin/rover ./cmd

run:
	go build -o bin/rover ./cmd && ./bin/rover

test:
	go test -v ./...

clean:
	rm bin/*
