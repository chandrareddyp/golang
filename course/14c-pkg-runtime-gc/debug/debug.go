package main

import (
	"fmt"
	"time"
)

func main() {
	<-  wait()
	fmt.Println("done")
}

func wait() <-chan struct{} {
	 done := make(chan struct{})

	 second := time.Tick(time.Second*1)
	 i:=0
	 for {
		select{
			case <- second:
				fmt.Println("Tick")
				i++
				if i>=20{
					close(done)
					return done
				}
		}
	}
}