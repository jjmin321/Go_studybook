package main

import "fmt"

func main() {
	//열거형
	//상수를 사용하는 일정한 규칙에따라 숫자를 계산 및 증가시키는 묶음
	const (
		Jan = iota
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)
}
