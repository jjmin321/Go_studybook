//29-1

package main

import (
	oper "GO_studybook/go_files/studybook/Arithmetic/arithmetic" //alias 사용(패키지 중복 또는 약자로 사용)
	"fmt"
	//"Arithmetic/arithmetic" //alias 미사용
)

func main() {
	//사용자 패키지 작성 & GO 문서화
	//기준은 GOPATH/src ("/Users/jejeongmin/Documents/go/src") //src 안에 폴더들이 더 있다면 src 아래 경로 다 타이핑해줘야함.
	//패키지 폴더명(디렉토리명) 명확하게 지정
	//패키지 파일의 package 이름으로 사용한다. -> 길면 alias 사용
	//package main 을 제외하고 package 문서에 등록
	//지금까지 우리는 패키지를 사용해왔다.
	//기본적으로 GOROOT의 패키지(pkg) 검색 하고 -> GOPATH의 패키지(src/pkg) 검색
	//go install 명령어 실행의 경우 -> GOPATH/pkg에 등록(문서화)
	//godoc -http=:6060(임의의 포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인(패키지, 타입, 메소드) 주석

	//패키지 사용 예제(사칙연산)
	nums := oper.Numbers{100, 10}
	fmt.Println(" 100 + 10 = ", nums.Plus())
	fmt.Println(" 100 - 10 = ", nums.Minus())
	fmt.Println(" 100 * 10 = ", nums.Multi())
	fmt.Println(" 100 / 10 = ", nums.Divide())
	fmt.Println(" (100 * 100) + (10 * 10) = ", nums.SquarePlus())
	fmt.Println(" (100 * 100) - (10 * 10) = ", nums.SquareMinus())
}
