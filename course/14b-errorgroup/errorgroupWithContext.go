package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	var wg sync.WaitGroup

	x := 1
	
    // Define tasks with potential errors
    wg.Add(1)
    eg.Go(func() error {
        defer wg.Done()
        // Simulate fetching user data
        if x == 0 { // Randomly fail for demonstration
            return errors.New("failed to fetch user data")
        }
        fmt.Println("Fetched user data")
        return nil
    })

    wg.Add(1)
    eg.Go(func() error {
        defer wg.Done()
        // Simulate retrieving shopping cart
        if x == 1 { // Randomly fail for demonstration
            return errors.New("failed to retrieve shopping cart")
        }
        fmt.Println("Retrieved shopping cart")
        return nil
    })

    wg.Add(1)
    eg.Go(func() error {
        defer wg.Done()
        // Simulate rendering the message
        time.Sleep(3 * time.Second) // Simulate some processing
        fmt.Println("Rendered welcome message")
        return nil
    })

	go func(){
		select {
		case  <- time.After(1 * time.Second):
			fmt.Println("Timeout")
			cancel()
		}
	}();
    // Wait for all tasks to finish (or timeout)
    err := eg.Wait()
	/*
    if err != nil {
        // Handle errors from any task or context timeout
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	*/
	if err != nil {
		fmt.Println("Final Error: ", err)
	}
}