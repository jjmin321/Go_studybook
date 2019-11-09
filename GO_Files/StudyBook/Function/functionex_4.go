//11-7

package main

import "fmt"

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
	//익명 함수(Anonymous Function)
	//즉시 실행

	//ex1
	func() {
		fmt.Println("ex1 : Annoymous Function") //매개변수 , 반환 값 없음
	}()

	//ex2
	msg := 5

	//ex3
	func(x, y int) {
		fmt.Println("ex3 : ", x+y) //매개변수 있음, 반환 값 없음 //주관적으로 ex2보다 많이 사용
	}(5, 6)

	//ex4
	r := func(x, y int) { //매개변수, 반환 값 있음
		fmt.Println(msg + x + y)
	}

	r(6, 5)

}
