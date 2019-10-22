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
	//Function
	//선언 : func 키워드로 선언
	//func 함수명(매개변수) (반환타입 or 반환 값 변수명)
	//func 함수명() (반환타입 or 반환 값 변수명)
	//func 함수명(매개변수)
	//func 함수명()
	//타 언어와 달리 return value 무한적으로 가능
	//변수에 함수 넣기 -> return 사용 시 변수에 value , 미사용시 변수는 function의 키가 됨.
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
