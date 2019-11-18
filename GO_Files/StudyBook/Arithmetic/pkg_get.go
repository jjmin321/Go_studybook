//사용자 패키지 설치 및 활용 예제
//29-2

package main

import (
	_ "fmt"

	_ "github.com/tealeg/xlsx"
)

func main() {
	//외부 저장소 패키지 설치
	//2가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 선언

}
