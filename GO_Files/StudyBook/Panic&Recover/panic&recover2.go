//25-2

package main

import "fmt"

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error message : ", s)
	}()

	panic("Error occurred!") //defer함수들만 실행하고 모든 코드 읽지않고 강제종료
	fmt.Println("panic Hi!") //호출되지 않음
}

func main() {
	//GO Recover 함수
	//에러 복구 가능
	//Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료되지 않는다.)
	//즉, 코드 흐름을 정상상태로 복구하는 기능
	//Panic에서 설정하는 메시지를 받아올 수 있다.

	runFunc()

	fmt.Println("Hello Golang!")
}
