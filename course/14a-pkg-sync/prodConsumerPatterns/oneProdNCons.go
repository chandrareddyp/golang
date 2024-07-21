package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main(){
	in := Produce()
	var wg sync.WaitGroup
	wg.Add(10)
	Consume(in, &wg, 10)
	wg.Wait()
}

func Produce() chan string{
	out := make(chan string)
	go func(){
		for i:=0;i<100;i++{
			time.Sleep(50 * time.Millisecond)
			out <- fmt.Sprintf("Message %d", i)
		}
		fmt.Println("producer setnt 100 messages, done.")
		close(out)
	}()
	return out
}

func Consume(ch chan string, wg *sync.WaitGroup, count int){
	time.Sleep(2 * time.Second)
	for i:=0;i<count;i++{
		str := fmt.Sprintf("consumer %d", i)
		go func(s *string){
			for msg := range ch{
				time.Sleep(10 * time.Millisecond)
				fmt.Println("recived message by consumber:",*s, " message:",msg)
				runtime.Gosched()
			}
			wg.Done()
		}(&str)
	}
}
