package main

import "fmt"

const (
	circumference = 3.14
)

type circle struct {
	radius float64
}

type calculate interface {
	calculate()
}

func (c circle) calculate() {
	fmt.Printf("출력\n반지름(Radius) : %0.3f\n넓이(Weight) : %0.3f", c.radius, c.radius*c.radius*circumference)
}

func main() {
	a := circle{1.5}
	a1 := calculate(a)
	a1.calculate()
}
