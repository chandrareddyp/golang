package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"unsafe"
)

func main() {
	//1. Get memory statistics
	 GCMemStat()

}

func GCMemStat() {
	fmt.Println("Hello, playground")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
 	// Print memory statistics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v MiB\n", memStats.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", memStats.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", memStats.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", memStats.NumGC)

	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
    // Allocate 1 GB of memory
    size := 1 * 1024 * 1024 * 1024 // 1 GB in bytes
    largeSlice := make([]byte, size)

    // Use the slice to avoid compiler optimizations that might remove it
    for i := range largeSlice {
        largeSlice[i] = byte(i % 256)
    }

	 
    // Print memory statistics
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Alloc = %v MiB\n", memStats.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MiB\n", memStats.TotalAlloc/1024/1024)
    fmt.Printf("Sys = %v MiB\n", memStats.Sys/1024/1024)
    fmt.Printf("NumGC = %v\n", memStats.NumGC)
 
	uintPtr := unsafe.Sizeof(largeSlice)

	fmt.Printf("Size of largeSlice = %v bytes\n", uintPtr)

	// Free up the memory
	largeSlice = nil
	
	runtime.ReadMemStats(&memStats)
	fmt.Printf("After freeup ******* Alloc = %v MiB\n", memStats.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MiB\n", memStats.TotalAlloc/1024/1024)
    fmt.Printf("Sys = %v MiB\n", memStats.Sys/1024/1024)
    fmt.Printf("NumGC = %v\n", memStats.NumGC)


	// Force GC to clear up the memory
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	fmt.Printf("After GC ******* Alloc = %v MiB\n", memStats.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MiB\n", memStats.TotalAlloc/1024/1024)
    fmt.Printf("Sys = %v MiB\n", memStats.Sys/1024/1024)
    fmt.Printf("NumGC = %v\n", memStats.NumGC)



	fmt.Println("done")
}