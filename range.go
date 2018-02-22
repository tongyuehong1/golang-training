package main

import "fmt"

func main() {
	num := []int{2, 3, 4}
	sum := 0
	for _, num := range num {
		sum += num
	}
	fmt.Printf("sum: %d\n", sum)
	for i, num := range num {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	for i := range num {
		num[i] = num[i] << 2
	}
	fmt.Println("输出是", num)
}
