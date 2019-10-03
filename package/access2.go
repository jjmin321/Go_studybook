package main

import (
	"fmt"
	alias "package/lib"
	_ "package/lib2"
)

func main() {
	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", alias.CheckNum(20))
}
