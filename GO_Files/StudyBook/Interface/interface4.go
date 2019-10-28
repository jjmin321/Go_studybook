package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (animal Dog) run() {
	fmt.Println(animal.name, "is running now!")
}

func (animal Cat) run() {
	fmt.Println(animal.name, "is running now!")
}

func act(animal interface{ run() }) { //은근 많이 쓰인다. 알고 있기를.
	animal.run()
}
func main() {

	//DuckTyping example
	//덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	//Go의 덕타이핑의 중요한 특징: 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다.

	//익명 인터페이스 사용 에제
	//ex1

	a := Dog{"poll", 10} // Dog타입의 변수를 선언함.
	skill(a)             //Dog타입에 맞는 메소드가 있다면 알아서 Dog타입 애들은 act인터페이스를 가짐. go의 어마무시한 덕타이핑 프로그래밍으로 인하여.
	b := Dog{"je", 85}
	skill(b)
	c := Cat{"dark", 15}
	skill(c)
	d := Cat{"white", 13}
	skill(d)

	//
	//b := Dog{"je", 85}
	//c := Cat{"dark", 15}
	//d := Cat{"white", 13}

	//animal1 := act(a) //animal1은 a의 행동을 나타내는 인터페이스
	//animal2 := act(b)
	//animal3 := act(c)
	//animal4 := act(d)
	//

	//
	// var animal1 act
	// var animal2 act
	// animal2 = Dog{"je", 85}
	// var animal3 act
	// animal3 = Cat{"dark", 15}
	// var animal4 act
	// animal4 = Cat{"white", 13}

	// skill(animal1)
	// skill(animal2)
	// skill(animal3)
	// skill(animal4)
	//
}
