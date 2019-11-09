//12-2

package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()
	func() {
		fmt.Println("sex6974")
	}()
}

func main() {
	//defer 함수 실행(지연)
	//defer를 호출한 함수가 종료되기 직전에 호출 된다.
	//타 언어의 Finally 문과 비슷
	//주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제
	//ex1
	sayHello("Golang!")
}
