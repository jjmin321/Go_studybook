package main

import (
	"fmt"
	sex "package/lib"
	_ "package/lib2"
)

func main() {
	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", sex.CheckNum(20))
}
