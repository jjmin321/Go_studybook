package main

import (
	"fmt"
	"os"
)

func main() {
	//package : code module and recycle
	//응집도, 결합도
	//GO : 패키지 단위의 독립이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고
	//패키지 이름 = 디렉토리 이름
	//같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다
	//네이밍 규칙 : 소문자 private , 대문자 public
	//GO : main package is recognize spacially -> 컴파일러 공유 라이브러리 X, program'sstart point

	var name string

	fmt.Println("이름은? : ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
