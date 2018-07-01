package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var a int32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt32(&a, 2) // 原子操作
			time.Sleep(time.Millisecond)
		}()
	}
	someone := atomic.LoadInt32(&a)
	time.Sleep(time.Second)
	fmt.Println("someone:", someone)
}
