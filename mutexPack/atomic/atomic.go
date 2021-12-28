package main

import(
	"fmt"
	"sync/atomic"
	"sync"
)

func AtomicCounterDemo() {
	var opt uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				//because opt is been accessed by multiple goroutines
				//unless otherwise atomic package use opt value will
				//be wrong
				atomic.AddUint64(&opt, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops", opt)
}

func main() {
	AtomicCounterDemo()
}
