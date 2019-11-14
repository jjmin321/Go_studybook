//23-3

package main

import (
	"errors"
	"fmt"
)

func main() {
	//에러 처리
	//errors 패키지의 New 메소드를 활용한 에러 생성
	//Errorf 보다 더 많이 사용

	var err1 error = errors.New("Error occurred - 1")
	err2 := errors.New("Error occurred - 2")

	fmt.Println("error1 : ", err1)
	fmt.Println("error2 : ", err2)
}
