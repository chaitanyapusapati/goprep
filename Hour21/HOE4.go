package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("ARCH :", runtime.GOARCH)
	fmt.Println("OS :", runtime.GOOS)
	fmt.Println("CPU :", runtime.NumCPU())
	fmt.Println("Start Goroutines :", runtime.NumGoroutine())
	count := 0
	const inc = 100
	wg.Add(inc)
	fmt.Println("Mid Goroutines :", runtime.NumGoroutine())
	var mu sync.Mutex
	for i := 0; i < inc; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goes Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("End Goroutines :", runtime.NumGoroutine())
	fmt.Println("Count :", count)

}
