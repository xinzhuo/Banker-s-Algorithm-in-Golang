#! /bin/bash
go build src/lifo.go
go build src/fifo.go
go build src/bankers.go

./lifo <samples/test1.txt >& liforesult/test1.out
./lifo <samples/test2.txt >& liforesult/test2.out
./lifo <samples/test3.txt >& liforesult/test3.out
./lifo <samples/test4.txt >& liforesult/test4.out
./lifo <samples/test5.txt >& liforesult/test5.out
./lifo <samples/test6.txt >& liforesult/test6.out
./lifo <samples/test7.txt >& liforesult/test7.out
./lifo <samples/test8.txt >& liforesult/test8.out
./lifo <samples/test9.txt >& liforesult/test9.out
./lifo <samples/test10.txt >& liforesult/test10.out

./fifo <samples/test1.txt >& fiforesult/test1.out
./fifo <samples/test2.txt >& fiforesult/test2.out
./fifo <samples/test3.txt >& fiforesult/test3.out
./fifo <samples/test4.txt >& fiforesult/test4.out
./fifo <samples/test5.txt >& fiforesult/test5.out
./fifo <samples/test6.txt >& fiforesult/test6.out
./fifo <samples/test7.txt >& fiforesult/test7.out
./fifo <samples/test8.txt >& fiforesult/test8.out
./fifo <samples/test9.txt >& fiforesult/test9.out
./fifo <samples/test10.txt >& fiforesult/test10.out

./bankers <samples/test1.txt >& bankersresult/test1.out
./bankers <samples/test2.txt >& bankersresult/test2.out
./bankers <samples/test3.txt >& bankersresult/test3.out
./bankers <samples/test4.txt >& bankersresult/test4.out
./bankers <samples/test5.txt >& bankersresult/test5.out
./bankers <samples/test6.txt >& bankersresult/test6.out
./bankers <samples/test7.txt >& bankersresult/test7.out
./bankers <samples/test8.txt >& bankersresult/test8.out
./bankers <samples/test9.txt >& bankersresult/test9.out
./bankers <samples/test10.txt >& bankersresult/test10.out
