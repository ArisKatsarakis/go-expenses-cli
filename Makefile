all:	build run 
build: 
	go build -o incexp ./cmd/incexp/main.go
run:
	./incexp
