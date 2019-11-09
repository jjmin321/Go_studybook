//4-1

package main

import "fmt"

func main() {
	//제어문(조건문) - switch
	//switch 뒤 표현식(Expression) 생략 가능
	//case 뒤 표현식(Expression) 사용 가능
	//자동 break 때문에 fallthrough 존재
	//Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

	//ex1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	//ex2
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	//ex3
	switch c := "go"; c {
	case "go":
		fmt.Println("GO!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음")
	}

	//ex4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("Golang")
	case "java":
		fmt.Println("JAVA")
	default:
		fmt.Println("NONE!")
	}

	//ex5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i is smaller than j")
	case i == j:
		fmt.Println("i is equal with j")
	case i > j:
		fmt.Println("i is bigger than j")
	}
}
