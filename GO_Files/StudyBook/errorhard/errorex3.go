//24-3

package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

//예외(에러) 처리 구조체
type PowError struct {
	time    time.Time   //에러발생 시간
	value   interface{} //파라미터
	message string      //에러 메세지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) { //error라는 인터페이스가 덕타이핑을 통해 Error()이 있으면 error라고 알아서 읽어줬으므로
	//Error()을 가지고 있는 PowError구조체가 error형으로 반환될 수 있게 해줍니다.
	//원래는 errors.New로 사용했지만 내가 정의한 PowError 구조체를 통해
	//time, value, message까지 내가 원하는 출력 값을 출력하게 할 수 있습니다.
	if f == 0 {
		return 0, PowError{time.Now(), f, "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) { //숫자를 넣은게 아니라 이상한 타입을 넣었을 때
		return 0, PowError{time.Now(), f, "숫자를 넣어주세요."}
	}
	if math.IsNaN(i) { //숫자를 넣은게 아니라 이상한 타입을 넣었을 때
		return 0, PowError{time.Now(), f, "숫자를 넣어주세요."}
	}
	return math.Pow(f, i), nil
}

func main() {
	//에러 처리 고급
	//error 타입이 아닌 경우 에러 처리 방법
	//Error 메소드를 구현해서 사용자 정의 에러 처리 예제 심화
	//구조체를 사용해서 세부적인 정보 출력
	val, err := Power(7, 3)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("val : ", val)
	val2, err2 := Power(0, 3)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	fmt.Println("val : ", val2)
}
