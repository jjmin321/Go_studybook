package main

import "fmt"

func main() {
	//제어문(조건문) - switch
	//switch 뒤 표현식(Expression) 생략 가능
	//case 뒤 표현식(Expression) 사용 가능
	//자동 break 때문에 fallthrough 존재
	//Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

	// 1.switch문과 비교 연산자를 사용하여 a의 값에 따라 출력값이 달라지는 프로그램의 빈칸을 채우세요.
	
	a := - 7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case 
		fmt.Println(a, "는 0")
    case 
		fmt.Println(a, "는 양수")
	}

	// 1.switch문을 사용하여 조건 c의 값에 따라 출력값이 달라지는 프로그램의 빈칸을 채우세요.

	switch c := "go"; c{
	case
		fmt.Println("GO!")
	case "java":
		fmt.Println("Java!")

		fmt.Println("일치하는 값 없음")
	}
}