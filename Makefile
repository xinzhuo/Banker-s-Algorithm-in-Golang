all: bankers fifo lifo

bankers: src/bankers.go
	go build src/bankers.go

fifo: src/fifo.go
	go build src/fifo.go

lifo: src/lifo.go
	go build src/lifo.go

exec: runner.sh
	./runner.sh

clean: 
	rm -rf *~ *.orig bankers fifo lifo
