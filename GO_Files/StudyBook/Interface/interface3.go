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

type act interface {
	run()
	jump()
}

func (animal Dog) run() {
	fmt.Println(animal.name, "is running now!")
}

func (animal Cat) run() {
	fmt.Println(animal.name, "is running now!")
}

func (animal Dog) jump() {
	fmt.Println(animal.name, "is jumping now!")
}

func (animal Cat) jump() {
	fmt.Println(animal.name, "is jumping now!")
}

func skill(a act) {
	a.run()
	a.jump()
}

func main() {
	a := Dog{"poll", 10}
	b := Dog{"je", 85}
	c := Cat{"dark", 15}
	d := Cat{"white", 13}

	animal1 := act(a) //animal1은 a의 행동을 나타내는 인터페이스
	animal2 := act(b)
	animal3 := act(c)
	animal4 := act(d)

	skill(animal1)
	skill(animal2)
	skill(animal3)
	skill(animal4)
}
