//24-1

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//f의 i제곱 구하는 함수
func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 입력하지 말아주세요.")
	}
	return math.Pow(f, i), nil //제곱, 에러없음 반환
}

func main() {
	//에러 처리 고급
	//Go error 패키지 New 메소드 사용 -> 사용자 에러 처리 생성

	//ex1
	if f, err := Power(7, 3); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(f)
	}
	//ex2
	if f, err := Power(0, 3); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(f)
	}
}
