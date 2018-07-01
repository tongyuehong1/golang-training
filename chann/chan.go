/*
 * Revision History:
 *     Initial: 2018/06/24        Tong Yuehong
 */

package main

import (
	"fmt"
	"time"
)

func Chan(msg chan string) {
	var count int
	go func() {
		for {
			count = count + 1
			fmt.Println(<-msg, count)
		}
	}()

	go func() {
		for {
			count = count + 1
			fmt.Println(<-msg, count)
		}
	}()
}

func main() {
	ch := make(chan string, 5)
	for i := 0; i<5; i ++ {
		ch <- "1"
	}

	Chan(ch)
	time.Sleep(1*time.Second)
}
