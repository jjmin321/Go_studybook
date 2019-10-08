package main

import "fmt"

func main() {
	const (
		A = iota + 0.75*2
		DEFAULT
		_
		GOLD
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)
}
