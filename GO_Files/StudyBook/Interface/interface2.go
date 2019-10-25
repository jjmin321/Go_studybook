package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

//bite메소드 구현
func (a Dog) bite() {
	fmt.Println(a.name, "bites!")
}

//동물의 행동 인터페이스 선언

type Behavior interface {
	bite()
}

func main() {
	//인터페이스 구현 예제

	//ex1
	dog1 := Dog{"poll", 10}   //dog1은 변수입니다.
	hotdog1 := Behavior(dog1) //hotdog1은 dog1의 인터페이스입니다.
	hotdog1.bite()            //use interface
	dog1.bite()               //use receiver method
	dog2 := Dog{"je", 85}

	//ex2
	slicedog := []Behavior{dog1, dog2}

	for i, v := range slicedog {
		fmt.Println(i, "번째", v)
	}

	//값 형태로 실행
	for _, v := range slicedog {
		v.bite()
	}
}
