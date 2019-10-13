// 이 디렉토리는 JAVA 코딩문제를 GO로 직접 번역시킨 파일들을 모아놓은 디렉토리입니다.
// 디렉토리 이름이 _로 시작하는 이유는 _는 GO언어에서 skip한다는 문법으로 skip 하셔도 된다는 말입니다 ^_^

/*
밑변과 높이 정보를 저장할 수 있는 Triangle클래스를 정의하자.(그 안에 적절한 생성자도 정의하자.)
밑변과 높이 정보를 변경할 수 있는 메소드와 삼각형의 넓이를 계산해서 반환하는 메소드도 함꼐 정의하자.
이 클래스의 활용의 예를 보이는 main메소드도 함께 작성해야 한다.
*/

package main

import "fmt"

type Wide struct { //밑변과 높이 정보를 저장할 수 있는 구조체 선언
	bottom, height float32
}

func main() {
	a := Wide{3, 4}         //밑변과 높이를 각각 3, 4 로 선언
	a.change()              //pointer receiver method(클래스 대신 GO에서 제공하는 기능)은 a.메소드명() 으로만 접근이 가능합니다.
	fmt.Print(a.triangle()) //value receiver method(클래스 대신 GO에서 제공하는 기능)또한 a.메소드명() 으로만 접근이 가능합니다.
}

func (a Wide) triangle() float32 { //(a Wide)는 value receiver로 a구조체의 데이터를 복사해서 사용 후 반환해줍니다.
	return a.bottom * a.height / 2
}

func (a *Wide) change() { //(a *Wide)는 pointer receiver로 구조체의 값을 변경할 수 있음 (pointer로 접근했을 때만 구조체의 값 변경 가능)
	a.bottom = 12
	a.height = 12
}
