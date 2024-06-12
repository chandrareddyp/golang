package main

import (
	"fmt"
	"sync"
)

/*
This program is a simple demonstration of how to use a RWMutex to protect a map from being read and written to concurrently.
*/
func main() {

	var m map[int]int
	m = make(map[int]int)
	rw := sync.RWMutex{}
	//m2 := map[int]int{1: 2, 3: 4}

	var wg sync.WaitGroup
	wg.Add(4)
	go func(m map[int]int){
		defer wg.Done()
		for i:=0;i<1000;i++{
			rw.Lock()
			m[i] = i
			rw.Unlock()
		}
	}(m)
	go func(m map[int]int){
		defer wg.Done()
		for i:=0;i<100000;i++{
			rw.RLock()
			_ = m[i]
			rw.RUnlock()
			//fmt.Println(x)
		}
	}(m)
	go func(m map[int]int){
		defer wg.Done()
		for i:=0;i<100000;i++{
			rw.RLock()
			_ = m[i]
			rw.RUnlock()
			//fmt.Println(x)
		}
	}(m)
	go func(m map[int]int){
		defer wg.Done()
		for i:=0;i<100000;i++{
			rw.RLock()
			_ = m[i]
			rw.RUnlock()
			//fmt.Println(x)
		}
	}(m)
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Done")
}