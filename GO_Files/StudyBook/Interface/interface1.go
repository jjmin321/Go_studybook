//16-1

package main

import "fmt"

type test interface { //빈 인터페이스

}

func main() {
	//인터페이스
	//객체의 동작을 표현, 골격
	//단순히 동작에 대한 방법만 표시
	//추상화 제공
	//인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능
	//GO언어를 유연하게 사용 가능
	//덕타이핑 : GO언어의 독창적인 특성
	//인터페이스 -> 자바(전략패턴, 템플릿메소드, 팩토리메소드, 어댑터패턴......)
	//디자인패턴 측면에서 Client 입장 -> 메소드의 구체적인 클래스를 몰라도 인터페이스 정의된 메소드를 사용하는 객체 보장
	//-> 클래스간의(Go로 따지면 구조체) 결합도 감소 -> 유지보수성 향상, 기능 추가의 용이성 -> 독립적인 프로그래밍 가능

	//ex1
	/*
		type 인터페이스명 interface{
			메소드1() 반환 값(타입 형)
			메소드1()	//반환 값 없는 경우
		}
	*/

	var a test
	fmt.Println(a) //구조체나 인터페이스나 비었으면 nil이 출력됨.

}
