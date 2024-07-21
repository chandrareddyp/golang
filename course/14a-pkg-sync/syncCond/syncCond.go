package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    var mu sync.Mutex
    var cond = sync.NewCond(&mu)
    var tasksCompleted int

    // Goroutine 1: Wait for all tasks to be completed
    go func() {
        mu.Lock()
        defer mu.Unlock()
        for tasksCompleted < 5 {
            fmt.Println("Goroutine 1 waiting for tasks to complete...")
            cond.Wait() // Wait on the condition variable with the mutex held
        }
        fmt.Println("Goroutine 1: All tasks completed!")
    }()

    // Simulate completion of tasks
    for i := 0; i < 5; i++ {
        go func(taskID int) {
            // Simulate some work
            time.Sleep(time.Second)
            fmt.Printf("Task %d completed.\n", taskID)

            // Signal completion
            mu.Lock()
            tasksCompleted++
            cond.Signal() // Signal one waiting goroutine (cond.Broadcast() for all)
            mu.Unlock()
        }(i)
    }

    // Wait for Goroutine 1 to finish
    time.Sleep(6 * time.Second)
}
