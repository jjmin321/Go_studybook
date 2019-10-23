package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (a Employee) Calculate() float64 {
	return a.salary + a.bonus
}

func (a Executives) Calculate() float64 {
	return a.salary + a.bonus + a.specialBonus
}

type Executives struct {
	Employee     //is a 관계 	//Executives is Employee
	specialBonus float64
}

func main() {
	//구조체 임베디드 메소드 오버라이딩 패턴

	//ex1
	//직원
	kim := Employee{"kim", 2000000, 150000}
	lee := Employee{"lee", 1500000, 300000}

	fmt.Println(int(kim.Calculate()))
	fmt.Println(int(lee.Calculate()))

	//임원
	park := Executives{Employee{"park", 3000000, 400000}, 1000000}
	fmt.Println(int(park.Calculate())) //오버라이딩 : 정확한 값 반환. park 타입에 맞는 Calculate() 메소드를 찾아 계산.
}
