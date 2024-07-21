package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)



func main() {
	CpuProfile()
}

func CpuProfile() {
	
		// Create the profile file
		f, err := os.Create("cpuProfile")
		if err != nil {
		  fmt.Println(err)
		  return
		}
		defer f.Close()
	  
		// Start CPU profiling
		if err := pprof.StartCPUProfile(f); err != nil {
		  fmt.Println("Error starting CPU profile:", err)
		  return
		}
		defer pprof.StopCPUProfile() // Ensure profiling stops even on errors
	  
		// Run your code to be profiled here (replace with your actual logic)
		fmt.Println("Profiling started. Run your code here.")
		time.Sleep(10 * time.Second) // Simulate some work (replace with your actual code)
	  
		// Optional: Graceful stop after some time (consider replacing with a proper signal handler)
		// go func() {
		//   time.Sleep(5 * time.Second)
		//   fmt.Println("Stopping profiling...")
		//   pprof.StopCPUProfile()
		// }()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				fmt.Println("Profiling started. Run your code here.")
			}
		}()
		wg.Wait()
	  
		fmt.Println("Profiling finished.")

	  
}