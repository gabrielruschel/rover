build:
	mkdir -p bin
	go build -o bin/rover ./cmd

run:
	go build -o bin/rover ./cmd && ./bin/rover

clean:
	rm bin/*
