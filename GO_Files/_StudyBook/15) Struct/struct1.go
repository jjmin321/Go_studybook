package main

import "fmt"

type Account struct {
	number   string  //계좌 번호
	balance  float64 //잔액
	interest float64 //이자
	total    float64
}

func (a *Account) Calculate() {
	a.total = a.balance + (a.balance * a.interest)
}

func main() {
	//Struct (구조체)
	//Go의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	//ex1
	kim := Account{"245-901", 10000000.0000, 0.015, 0}
	lee := Account{"245-902", 12000000.0000, 0.015, 0}

	kim.Calculate()
	lee.Calculate()
	fmt.Printf("kim's balance : %f, kim's total : %f\n", kim.balance, kim.total)
	fmt.Printf("lee's balance : %f, lee's total : %f\n", lee.balance, lee.total)
}
