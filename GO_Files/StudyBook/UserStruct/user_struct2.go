//14-2

package main

import "fmt"

type cnt int

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

	//기본 자료형 사용자 정의 타입

	//ex1

	a := cnt(5) //이렇게는 쓰지 않음
	fmt.Println(a)

	//ex2

	var b cnt = 5 //이렇게 사용함
	fmt.Println(b)

	//ex3

	//testConverT(b)	//error	예외 발생(중요!) 사용자 정의 타입 <-> 기본 타입 : 매개 변수 전달 시에 변환해야 사용 가능(int(변수))
	testConverT(int(b))
	testConverD(b)
}

func testConverT(i int) {
	fmt.Println("ex3 : (Default type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex3 : (custom type) : ", i)
}
