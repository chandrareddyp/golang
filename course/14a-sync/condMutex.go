package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
var mx sync.Mutex
cond := sync.NewCond(&mx)
var taskDone int
wg := sync.WaitGroup{}
go func(){
	wg.Add(1)
	cond.L.Lock()
	defer cond.L.Unlock()
	for taskDone < 5 {
		cond.Wait()
		fmt.Println("Task Done: ", taskDone)
	}
	wg.Done()
}()
for i := 0; i < 5; i++ {
	go func(taskId int){
		wg.Add(1)
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		taskDone++
		fmt.Println("Task just Done: ", taskDone)
		cond.Signal()
		cond.L.Unlock()
		wg.Done()
	}(i)
}
wg.Wait()
//time.Sleep(10 * time.Second)

}