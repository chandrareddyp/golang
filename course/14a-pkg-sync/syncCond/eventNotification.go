package main

import (
	"fmt"
	"sync"
	"time"
)

const maxWorkersCount = 10

func main() {
	var counter int32

	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	wg.Add(maxWorkersCount)

	for i :=0; i<  maxWorkersCount;i++ {
		go func(workerID int) {
			defer wg.Done()

			fmt.Printf("Worker %d performing work\n", workerID)
			time.Sleep(1 * time.Second) // Simulate work

			cond.L.Lock()
			defer cond.L.Unlock()

			counter++

			if counter == maxWorkersCount {
				fmt.Println("All workers have reached the barrier")
				cond.Broadcast()
			} else {
				fmt.Printf("Worker %d is waiting at the barrier\n", workerID)
				cond.Wait()
			}

			fmt.Printf("Worker %d passed the barrier\n", workerID)
		}(i)
	}

	wg.Wait()
}