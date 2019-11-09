//15-9

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

type Executives struct {
	Employee     //is a 관계 	//Executives is Employee
	specialBonus float64
}

func main() {
	//구조체 임베디드 패턴
	//다른 관점으로 메소드를 재사용하는 장점 제공
	//상속을 허용하지 않는 Go언어에서 메소드 상속 활용을 위한 패턴

	//ex1
	//직원
	kim := Employee{"kim", 2000000, 150000}
	lee := Employee{"lee", 1500000, 300000}

	fmt.Println(int(kim.Calculate()))
	fmt.Println(int(lee.Calculate()))

	//임원
	park := Executives{Employee{"park", 3000000, 400000}, 1000000}
	fmt.Println(int(park.Employee.Calculate() + park.specialBonus))
	//fmt.Println(int(park.Calculate() + park.specialBonus))
}
