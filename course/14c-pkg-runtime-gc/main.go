package main

import (
	"fmt"
	"runtime"
	"time"
	// "gopractice.com/runtime/goshed"
)

/*
How to run:
go run code.go goschedFunc.go

OR
go mod init gopractice.com/runtime
go mod tidy
go build

*/

func main() {
	//1. Get memory statistics
	 GCMemStat()

	//2. Go Sched
	// GoSched22();
}

func GoSched22() {
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 1:", i)
            runtime.Gosched() // Yield the processor
        }
    }()

    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 2:", i)
            runtime.Gosched() // Yield the processor
        }
    }()

    time.Sleep(time.Second) // Wait for goroutines to complete
}