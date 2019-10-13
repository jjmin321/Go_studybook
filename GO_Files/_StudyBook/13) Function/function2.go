package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println(a + b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i * 10
}

func main() {
	//함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)

	//ex1
	sum(100, add) //함수(콜백) 전달

	//ex2
	//call by value
	a := 100
	multi_value(a)
	fmt.Println(a)

	//ex3
	//call by reference
	b := 100
	multi_reference(&b)
	fmt.Println(b)
}
