package main

import (
	"fmt"
	"time"
)




func main() {
    ch := make(chan int, 1) // Set buffer size to 1 for non-blocking behavior

    done := make(chan struct{})

    // Producer Goroutine
    go producer(ch, done)

    // Consumer Goroutine
    go consumer(ch, done)

    time.Sleep(1 * time.Second)

	close(done) // Signal to terminate goroutines
    time.Sleep(1 * time.Second)
    close(ch) // Close the channel to signal completion
   
    fmt.Println("Closed channel")
}

func producer(ch chan int, done chan struct{}) {
    i := 0
    for {
        select {
        case ch <- i:
			fmt.Println("Sent:", i)
            i++
        case <-done:
            fmt.Println("Producer exiting")
            return
        }
    }
}

func consumer(ch chan int, done chan struct{}) {
    for {
        select {
        case x, ok := <-ch:
            if !ok {
                fmt.Println("Channel closed")
                return
            }
            fmt.Println("Received:", x)
        case <-done:
            fmt.Println("Consumer exiting")
            return
        }
    }
}
