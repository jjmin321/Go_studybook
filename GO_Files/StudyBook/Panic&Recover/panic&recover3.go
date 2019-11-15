//25-3

package main

import "fmt"

func runFunc() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error message : ", s)
		}
		//프로그램 내부에서 일어난 또는 내가 만든 에러가 있으면 recover이 그 에러를 갖고 있음(에러는 defer실행 후 자기가 있는 함수 종료하므로 defer 안에 선언!).
	}()

	a := [3]int{1, 2, 3}
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
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
