package main

import "fmt"

type some struct {
	big, small int
}

var m map[string]some

func main() {
	m = make(map[string]some)
	m["first second"] = some{
		1, 2,
	}
	m["a"] = some{3, 4}
	var n = map[string]some{
		"a b": {66, 88},
		"h k": {666, 888},
		"g":   {55, 99},
	}
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(m["first second"])

	// 不用函数体
	var student map[string]string
	student = make(map[string]string)
	student["Li"] = "male"
	student["Susan"] = "female"
	fmt.Println(student)
	for name := range student {
		fmt.Println(name, "is", student[name])
	}
	sex, ok := student["Lucy"]
	if ok {
		fmt.Println("Lucy is", sex)
	} else {
		fmt.Println("not exits")
	}
	// 增删改
	student["Lucy"] = "female"
	fmt.Println("Lucy is", student["Lucy"])
	delete(student, "Li")
	fmt.Println(student["Li"])
	student["Lucy"] = "none"
	fmt.Println(student)
	sex, ok = student["Lucy"]
	fmt.Println("xiu xiu xiu", sex, "Present?", ok)
}
