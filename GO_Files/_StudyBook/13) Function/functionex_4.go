package main

import "fmt"

func main() {
	//함수 고급
	//익명 함수(Anonymous Function)
	//즉시 실행

	//ex1
	func() {
		fmt.Println("ex1 : Annoymous Function") //매개변수 , 반환 값 없음
	}()

	//ex2
	msg := "Hello Golang!"

	func(m string) {
		fmt.Println("ex2 : ", m) //매개변수 있음, 반환 값 없음
	}(msg)

	//ex3
	func(x, y int) {
		fmt.Println("ex3 : ", x+y) //매개변수 있음, 반환 값 없음 //주관적으로 ex2보다 많이 사용
	}(5, 6)

	//ex4
	r := func(x, y int) int { //매개변수, 반환 값 있음
		return x * y
	}(10, 10)

	fmt.Println("ex4 : ", r)

}
