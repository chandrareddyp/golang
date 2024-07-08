package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main(){

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// A goroutine to wait for the signal.
	go func(){
		val := <- ch
		println("Signal received:", val)
		os.Exit(0)
	}()
	// Block until a signal is received.
	//sig := <-ch
	//fmt.Println("Signal received:", sig)

	select{}

}