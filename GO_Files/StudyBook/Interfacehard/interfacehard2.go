package main

import (
	"fmt"
)

func main() {
	//타입 현황
	//Type Assertion
	//실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	//인터페이스(타입) -> 형 변환
	//interfaceVal.(type)

	//ex1
	var a interface{} = 15
	fmt.Printf("%T, %#v", a, a)

}
