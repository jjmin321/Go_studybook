// 함수 심화(3)
package main

import "fmt"

func prtHello(n int) {
	if n == 1 {
		fmt.Println("Hello")
		return
	}
	fmt.Println("Hello")
	prtHello(n - 1)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)

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
	//함수 고급
	//재귀 함수(Recursion)
	//프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이 : 장점
	//코드 이해하기 어렵고, 기억공간을 많이 사용, 무한루프 가능성

	//ex1
	x := fact(5)
	fmt.Println(x)

	//ex2
	prtHello(5)

}
