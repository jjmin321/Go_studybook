//23-1

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//에러 처리
	//에러 처리 : 소프트웨어의 품질 향상 가장 중요한 것! -> 유형코드 및 에러 정보 등을 정보를 남기는 것!
	//Go에서는 기본적으로 error 패키지에서 error 인터페이스를 제공
	//type error interface { Error() string }
	//즉, Error 메소드를 구현하면 사용자 정의 에러 처리 제작 가능
	//기본적으로 메소드마다 리턴 타입 2개(리턴값, error)
	//주로 Errorf(에러 타입 리턴) 메소드, Fatal(프로그램 종료) 메소드를 통해서 에러 출력

	//기본적인 메소드 에러 처리 예제
	f, err := os.Open("unnamedfile") //err
	if err != nil {
		log.Fatal(err.Error()) //방법 1
		// log.Fatal(err) //방법2
	}
	fmt.Println("===============") //log.Fatal로 종료돼서 실행되지 않음.
	fmt.Println(f.Name())          //log.Fatal로 종료돼서 실행되지 않음.
}
