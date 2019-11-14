//23-2

package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) { //메소드 리턴 값 error 타입 중요!
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", fmt.Errorf("%d을 입력했습니다. 에러 발생!", n) //Errorf 잘 안 써요~ 23-4파트가 중요함!
}

func main() {
	//에러 처리
	//Errorf를 이용한 에러 처리

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)

	b, err := notZero(0) //0넣고 테스트
	if err != nil {
		log.Fatal(err)
	}
	//Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println(b)
	fmt.Println("End error Handling!")
}
