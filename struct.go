package main
import "fmt"

type student struct {
	name string
	sex string
	major string
	id int
}
type books struct {
	title string
	id int
	author string
}
func main() {
	var first student
	first.name = "first"
	first.sex = "male"
	first.major = "computer"
	first.id = 1
	second := student{"second", "female", "math", 2}
	fmt.Printf("second的姓名，性别，专业，学号分别是： %s, %s, %s, %d\n", second.name, second.sex, second.major, second.id)
	fmt.Printf("first的姓名，性别，专业，学号分别是： %s, %s, %s, %d\n", first.name, first.sex, first.major, first.id)
	
	something := books{"计算机基础", 666, "不知道谁"}
	showbook(something)

	another := books{"计算机系统", 888, "还是不知道"}
	bookspointer(&another)


	example := []struct {
		computer string
		num int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	fmt.Println("computer的example(struct)是：", example)
}

func showbook(book books) {
	fmt.Printf("这本书的名字，编号，作者分别是：%s, %d, %s\n", book.title, book.id, book.author)
}

// 结构体指针
func bookspointer(book *books) {
	fmt.Printf("另一本书的名字，编号，作者分别是： %s, %d, %s\n", book.title, book.id, book.author)
}
