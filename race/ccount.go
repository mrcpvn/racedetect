package main

import (
	"fmt"
	"sync"
)

func main() {
	s := run()
	fmt.Printf("concurrent count = %v\n", s)
}

func run() int {
	//var mux sync.Mutex
	var wg sync.WaitGroup
	race := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//mux.Lock()
			race++
			//mux.Unlock()
		}()
	}
	wg.Wait()
	return race
}
