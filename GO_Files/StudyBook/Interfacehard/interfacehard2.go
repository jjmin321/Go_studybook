package main

import (
	"fmt"
	"reflect"
)

func main() {
	//타입 현황
	//Type Assertion
	//실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	//인터페이스(타입) -> 형 변환	//interfaceVal.(type)

	//ex1
	var a interface{} = 15
	b := a
	c := a.(int)
	// d := a.(float64) //안 됨

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))

	//ex2(저장된 실제 타입 검사)

	if v, ok := a.(float64); ok { // 실행 안됨. 왜냐하면 a는 애초에 int라서 ㅋㅋ
		fmt.Println(v, ok)
	}
	if v, ok := a.(int); ok { // v에 a.(int)를 넣고 ok에 a의 원본 타입이 int가 맞으면 참, 아니라면 거짓을 넣는다.
		fmt.Println(v, ok)
	}

}
