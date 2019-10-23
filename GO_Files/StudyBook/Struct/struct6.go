package main

import "fmt"

type Car struct { //대문자로 선언 - 외부에서도 사용 가능
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec
}

type spec struct { //소문자로 선언 - 내부에서만 사용 가능
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	//중첩 구조체

	//ex1
	car1 := Car{"520d", "purple", "samsung", spec{40, 10, 400}}
	fmt.Println(car1)         //car1의 구조체값을 보여줌
	fmt.Printf("%#v\n", car1) //car1의 모든 것을 보여줌

	//ex2
	fmt.Println(car1.detail.length)
	fmt.Println(car1.detail.height)
	fmt.Println(car1.detail.width)

}
