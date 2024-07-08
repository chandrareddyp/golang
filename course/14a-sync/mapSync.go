package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.RWMutex
    cond := sync.NewCond(&m)
	c := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&m, &c, &wg, cond)
		cond.Broadcast()
	}
	wg.Wait()
}

func worker(m *sync.RWMutex, c *int, wg *sync.WaitGroup, cond *sync.Cond){
	//m.Lock()
	//defer m.Unlock()
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for i := 0; i < 100; i++ {
		*c++
		if i == 5{
			fmt.Println("going to wait")
			cond.Wait()
			fmt.Println("wait done")
		}
		fmt.Println("Worker: ", *c)
	}
}