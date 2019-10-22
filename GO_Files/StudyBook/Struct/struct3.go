package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {
	//Struct (구조체)
	//Go의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	//ex1
	c1 := car{"red", "220d"} //포인터형 X

	c2 := new(car) //포인터형
	c2.color, c2.name = "white", "sonata"

	c3 := &car{"black", "520d"} //포인터형

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}
