package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main(){

	ctx, cancel := context.WithCancel(context.Background())
	
	task := func(ctx context.Context){
		
		for{
			select {
			case <- ctx.Done():
				println("Task is done, context is cancelled.")
				os.Exit(0)
			default:
				time.Sleep(5 * time.Second)
				ctx.Done()
				fmt.Println("Task is done. timeout")
				os.Exit(0)
			}
		}
	}
	go task(ctx)

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		val :=  <- ch
		println("Signal receiveddd:", val)
		cancel()
	}()

	select{}

}