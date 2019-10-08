package main

import (
	"fmt"
)

func main() {
	//Bool(Boolean)
	//조건부 논리 연산자랑 주로 사용 : !, &&, ||
	//암묵적 형변환X : 0, Nil -> False 변환 없음

	//ex1
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex2 : ", b2)
	fmt.Println("ex3 : ", b3)

	fmt.Println("ex4 : ", b3 == b3)
	fmt.Println("ex5 : ", b1 && b3)
	fmt.Println("ex6 : ", b1 || b2)
	fmt.Println("ex7 : ", !b1)

}
