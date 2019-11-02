package main

import "fmt"

func checkType(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Println("This is bool,", a)
	case int, int8, int16, int32, int64:
		fmt.Println("This is int,", a)
	case float32, float64:
		fmt.Println("This is float,", a)
	case nil:
		fmt.Println("This is nil,", a)
	case string:
		fmt.Println("This is string,", a)
	default:
		fmt.Println("This is not bool, int, float, nil, string!,", a)
	}
}

func main() {
	//실제 타입 검사 switch 주로 사용
	//빈 인터페이스는 어떠한 자료형도 전달받을 수 있으므로 타입 체크를 통해 변환 후 사용 가능

	//ex1
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("Hello World!")
}
