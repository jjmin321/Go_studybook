//15-2

package main

import "fmt"

type Account struct {
	number   string  //계좌 번호
	balance  float64 //잔액
	interest float64 //이자
}

func main() {
	//Struct (구조체)
	//Go의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)
	//다양한 선언 방법
	//&struct, struct : &struct 포인터를 받아오기 역참조를 또 하기 때문에 속도는 조금 느리다.
	//인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메소드에 포인터 리시버를 사용할 경우 반드시 &struct 사용해야함.

	//선언 방법 1	//주관적으로 거의 안 씀. 인터페이스 사용 시 강제적으로 이렇게 사용해야함.
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방법 2	//주관적으로 가장 많이 사용함. 인터페이스 사용 시 강제적으로 이렇게 사용해야함.
	hong := &Account{"245-902", 15000000, 0.04}

	//선언 방법 3	//주관적으로 거의 안 씀. 인터페이스 사용 시 강제적으로 이렇게 사용해야함.
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex2 : ", hong)
	fmt.Println("ex3 : ", lee)

	fmt.Printf("ex1 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex3 : %#v\n", lee)
}
