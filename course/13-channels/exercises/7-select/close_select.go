package main

import (
	"fmt"
	"time"
)

func main() {
    ch := make(chan int)
    done := make(chan int)

    go func() {
        for i := 1; i < 5; i++ {
            ch <- i
			time.Sleep(1 * time.Second)
        }
        close(ch)
    }()

    go func() {
        // Simulate some other condition to stop the loop
        time.Sleep(2 * time.Second)
		fmt.Println("Closing done channel")
        close(done)
    }()

    for {
        select {
        case value, ok := <-ch:
            if !ok {
                fmt.Println("Channel is closed")
                return
            }
            fmt.Println("Received value:", value)
        case v, k := <-done:
            fmt.Println("Done signal received, exiting loop")
			fmt.Printf("Value: %v, Key: %v\n", v, k)
            return
        default:
            fmt.Println("No value ready, performing other work")
            time.Sleep(1 * time.Second)
        }
    }
}
