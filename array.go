package main

import "fmt"

// func main() {
// 	var n [6]int
// 	var i int
// 	for i = 0; i < 6; i++ {
// 		n[i] = i + 200
// 		fmt.Printf("n[%d]: %d\n", i, n[i])
// 	}
// }

// 向函数传递数组
func main () {
	var n = []int{2, 4, 6 ,8}
	var avg float32
	avg = getavg( n, 4 )
	fmt.Printf("avg是：%f\n", avg)
	
	// 多维数组
	var m = [2][3]int{{1,2,8}, {3,4,9}}
	var a,b int
	for a = 0; a < 2; a++ {
		for b = 0; b < 3; b++ {
			fmt.Printf("m[%d][%d]: %d\n", a , b , m[a][b])
		}
	}
	fmt.Println("m数组是：", m)
}
func getavg(arr []int, size int) float32{
	var i int
	var sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum/size)
	return avg
}
