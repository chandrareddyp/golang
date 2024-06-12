package main

import (
	"fmt"
	"time"
)

func countReports1(numSentCh chan int) int {
	// ?
	output := 0
	_, ok := <-numSentCh
	if !ok {
		fmt.Println("Channel is closed")
		return output
	}
	for numReports := range numSentCh {
		fmt.Printf("Received %v reports\n", numReports)
		output += numReports
	}
	return output
}

func countReports(numSentCh chan int) int {
	output := 0

	for {
		numReports, ok := <-numSentCh
		if !ok {
			fmt.Println("Channel is closed")
			return output
		}
		fmt.Printf("Received %v reports\n", numReports)
		output += numReports
	}
}

func countReports2(numSentCh chan int) int {
	output := 0
	fmt.Println("Starting....")
	for {
		select {
		case numReports:= <-numSentCh:
			 
			fmt.Printf("Received %v reports\n", numReports)
			output += numReports
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("Exiting....")
	return output
}

// TEST SUITE - Don't touch below this line

func test(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports2(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func main() {
	test(3)
	//test(4)
	//test(5)
	//test(6)
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Second * 1)
	}
	close(ch)
}
