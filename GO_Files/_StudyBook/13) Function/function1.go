//Function

package main

import (
	"fmt"
	"strconv"
)

//함수 위치는 어디에서나 가능
func helloworld() {
	fmt.Println("Hello world!")
}

func prtstr(a string) {
	fmt.Println(a)
}

func sum(x int, y int) int {
	return x + y
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

	//ex1
	helloworld()

	//ex2
	prtstr("Hello world!")

	//ex3
	result := sum(3, 4)
	fmt.Println(result)
	fmt.Println(strconv.Itoa(result))
}
