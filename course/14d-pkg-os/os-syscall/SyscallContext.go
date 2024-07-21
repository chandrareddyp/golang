package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)



func main() {
	//1. demo os/signal and syscall
	// Syscall()

	//2. demo os/signal and Context
	Context()
	//proc, err := os.StartProcess("ls", []string{"ls", "-l"}, &os.ProcAttr{})
	//proc.Release()
}


func Syscall(){
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	// signal.Notify(interrupt, syscall.SIGINT) // Register for SIGINT (Ctrl+C)

	go func() {
		tick := time.Tick(1 * time.Second)
		for{
			select{
			case <- tick:
				println("tick1111")
			case <- interrupt:
				println("interrupt1111")
				os.Exit(0)
				return
			}
		}
	}()
	go func() {
		tick := time.Tick(1 * time.Second)
		f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0644)
		for{
			select{
				case <- tick:
					f.Write([]byte("tick2222"))

				case <- interrupt:
					f.Write([]byte("interrupt2222"))
					println("interrupt2222")
					os.Exit(0)
					return
			}
		}
	}()
	select{}
}

func Context(){
	ctx, cancel := context.WithCancel(context.Background())
	//context.WithDeadline(ctx, time.Now().Add(5*time.Second))

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		tick := time.Tick(1 * time.Second)
		
		defer wg.Done()
		for{
			select{
			case <- tick:
				println("tick1111")
			case <- interrupt:
				println("interrupt1111")
				cancel()
				return
			}
		}
	}()
	wg.Add(1)
	go contextCancel(ctx, &wg)
	wg.Wait()
}

 func contextCancel(ctx context.Context, wg *sync.WaitGroup) {	
	defer wg.Done()
	tick := time.Tick(1 * time.Second)
	for{
		select{
			case <- tick:
				println("tick2222")
			case <- ctx.Done():
				println("interrupt2222")
				return
		}
	}
}

/*
Below code not working:::

func main() {
    ctx, cancel := context.WithCancel(context.Background())
   
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Listen for SIGINT (Ctrl+C) and SIGTERM
	go func() {
        task(ctx)
    }()
    go func() {
        sig := <-sigs // Wait for a signal
        fmt.Println("Received signal:", sig)
        cancel() // Cancel the context upon receiving a signal
    }()
    fmt.Println("Press Ctrl+C or send SIGTERM to terminate...")
    select {} // Block main routine until context is canceled or a signal is received
}
func task(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Task canceled")
            return
        default:
            fmt.Println("Task running...")
            time.Sleep(time.Second * 1)
        }
    }
}
*/