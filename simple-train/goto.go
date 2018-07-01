package main

import "fmt"

func main() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

// func main() {
//     for true  {
//         fmt.Printf("这是无限循环。\n");
//     }
// }
