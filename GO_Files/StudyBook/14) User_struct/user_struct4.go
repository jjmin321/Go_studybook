package main

import "fmt"

type shoppingBasket struct {
	cnt   int
	price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

//원본 수정(참조 형식)
func (b *shoppingBasket) rePurchase(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

/* 원본 수정 X (값 전달 방식)
func (b *shoppingBasket) rePurchase(cnt, price int) {
	b.cnt += cnt
	b.price += price
}
*/

func main() {
	//GO -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
	//객체 지향 -> 클래스(속성 : 멤버변수, 기능(상태 : 메소드)) : 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그래밍
	//클래스가 없는데 객체 지향 언어일까? -> 맞다
	//GO는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체 -> 클래스형태의 코딩 가능
	//객체지향의 기본 개념 -> GO에서 포함하고 있다. -> 객체 지향 프로그래밍 언어
	//상태, 메소드를 분리해서 정의(결합성 없음)
	//사용자 정의 타입 : 구조체, 인터페이스, 기본 타입(int, float, string...) , 함수
	//구조체와 -> 메소드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능하다 (객체지향이 맞다)
	//클래스는 구조체다. 구조체는 클래스가 아니다. 클래스는 구조체와 메소드의 연결과 같다.

	//리시버 메소드 전달(값, 참조) 형식
	//함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정 X) -> 맵, 슬라이스 등은 참조 전달(자동 포인터형 느낌)
	//리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능

	//ex1
	bs1 := shoppingBasket{5, 1000}
	fmt.Println("bs1 : ", bs1)
	fmt.Println(bs1.purchase())
	bs1.rePurchase(5, 1000)
	fmt.Println(bs1.purchase())

}
