package main

import (
	"fmt"
	"runtime"
	"time"
)



func goSched() {
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