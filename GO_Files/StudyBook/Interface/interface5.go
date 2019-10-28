package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1", s)
}

func main() {
	//인터페이스 활용(빈 인터페이스)
	//함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능) -> 모든 타입 지정 가능
	//빈 인터페이스로 모든 매개변수를 대신해서 쓸 수 있다.

	//ex1
	dog := Dog{"poll", 10}
	cat := Cat{"black", 15}

	printValue(dog)
	printValue(cat)
	printValue("Hi")
	printValue(15)
	printValue(15.5)
	printValue([3]int)
	printValue([3]int{})
	printValue(map)

}
