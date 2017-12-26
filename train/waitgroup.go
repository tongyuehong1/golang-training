package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

func init() {
	fmt.Println("Current Go Version:", runtime.Version())
}
func main() {
	runtime.GOMAXPROCS(1)

	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count * 2)
	for i := 0; i < count; i++ {
		//time.Sleep(200)
		go func() {
			fmt.Println("[%d]", i)
			wg.Done()
		}()
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Println("-%d-", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}