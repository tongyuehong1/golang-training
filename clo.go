package main

import "fmt"

// 匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
	 }
}

func add() int{
  i := 0
  i+=1
  // fmt.Println(i)
  return i
}
func main(){
   nextNumber := getSequence()

   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}