package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	// runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())
	// withoutWait()
	// withWait()
	writeWaitMute()
}

func withoutWait(){
	for i := 0; i < 5; i++ {
		go fmt.Println(i + 1)
	}
	fmt.Println("exit")
}

func withWait (){
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i + 1)
			wg.Done()
		}(i)

	}

	wg.Wait()
	fmt.Println("exit")
}

func writeWaitMute(){
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}