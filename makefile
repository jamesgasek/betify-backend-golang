default:
	go build -o bin/server ./src/server/server.go
	go build -o bin/client ./src/tester/tester.go
clean:
	rm -rf bin