package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

/*

problem: we need to read a file which has city and temparature data

assuming its well formeted

we need max, min, avg temparature of each city

read file
parse each line
no. of occurance for a city.
	for each need to maintain min and max, total temparature, and occurance
	caluclate avg.






*/

type City struct {
	// Name string
	Max float64
	Min float64
	Avg float64
	count int
}

func main() {
	
	file, err :=  os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	cityInfo := make(map[string]City)
	syncMap := sync.Map{}
	syncMap.Store("cityInfo", cityInfo)
	syncMap.Load("cityInfo")

	// no.of go routines

	// read 1k line at a time
	
	mapLock := sync.Mutex{}


	data := make([]byte, 1024)
	// dallas: 100
	// new york: 95.4
	// temp := new york: 95

	// calculate size of file 
	// divide by 1m = 1024 
	no. of go routines = 1024 // n-2
	n-2 maps 
	combile 
	runtime.NumCPU()
	// n, n-2

	wg := sync.WaitGroup{}
	
	wg.wait()

	channels - 
		no need of local maps 
		we just global maps - contentions not much, 
		consumber go routi9ne one 
		16 go func(ch chan int) {
			for {
				select {
				case <- ch:
					// read 1k line
					// process
				}
			}
			
		}()





	// 1mb file
	position :=0

	n, err := file.ReadAt(data, int64(position))
	if err != nil {
		fmt.Println("Error reading file")
		return
	}


	if err != nil {
		fmt.Println("Error reading file")
		return
	}


	rows := strings.Split(string(data), "\n")
	for _, row := range rows {
		// validation for empty row, well formated row
		vals := strings.Split(row, ":");
		if len(vals) == 0 {
			continue
		}
		if len(vals) != 2 {
			// log error 
			continue
		}
		city := vals[0]
		temp := vals[1]
		// validate temp
		// convert temp to float
		// check if city already exists
		city = strings.TrimSpace(city)
		string to float
		t := strconv.ParseFloat(temp, 64)

		if _, ok := cityInfo[city]; !ok {
			mapLock.Lock()
			cityInfo[city] = City{ Max: t, Min: t, Avg: t, count: 1}
			mapLock.Unlock()
		}else{
			mapLock.Lock()
			cityInfo[city].Avg = (cityInfo[city].Avg * cityInfo[city].count + t)/(cityInfo[city].count+1)
			cityInfo[city].count++
			if t > cityInfo[city].Max {
				cityInfo[city].Max = t
			}
			if t < cityInfo[city].Min {
				cityInfo[city].Min = t
			}
			mapLock.Unlock()
		}
	}
}