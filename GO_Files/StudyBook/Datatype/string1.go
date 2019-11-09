// 1-1

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//문자열
	//큰 따옴표 "", 백스쿼트 ``
	//Golang : 문자 char 타입 없음 -> rune(int32)로 문자 코드 값으로 표현
	//문자 : ''로 작성
	//자주 사용하는 escape : \\, \', \", \a(콘솔벨) , \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)

	//ex1
	var str1 string = "c:\\users\\user\\go\\src" // -> c:\users\user\go\src
	str2 := `c:\users\user\go\src`

	fmt.Println(str1)
	fmt.Println(str2)

	//ex2
	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요"

	fmt.Println(str3)
	fmt.Println(str4)

	//ex3
	//길이(바이트 수)
	fmt.Println("Hello, world! : ", len(str3))
	fmt.Println("안녕하세요 : ", len(str4))

	//ex4
	//길이(실제 길이))
	fmt.Println("Hello, world! : ", utf8.RuneCountInString(str3))
	fmt.Println("안녕하세요 : ", utf8.RuneCountInString(str4))
	fmt.Println("안녕하세요 : ", len([]rune(str4))) //Len으로 실제 길이 구하는 법
}
