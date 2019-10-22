package main

import "fmt"

func main() {
	// 1.출력값이 무엇이 될 지 생각해보고 주석을 사용하여 답을 맞춰보세요.
	// 2.boolean형 변수 a, b, c, d, e, f 의 값을 초기화하여 출력되는 값이 답과 같게 만드세요.
	fmt.Println("ex1 : ", true && true)
	fmt.Println("ex2 : ", true && false)
	fmt.Println("ex3 : ", true || false)
	fmt.Println("ex4 : ", false || false)
	fmt.Println("ex5 : ", !true)
	fmt.Println("ex6 : ", !false)

	var a bool
	var b bool
	var c bool
	var d bool
	var e bool
	var f bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
