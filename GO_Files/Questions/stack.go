package main

import "fmt"

func main() {
	MAX := 10
	stack := [10]int{}
	var top int

	init_stack := func() {
		top = -1
	}

	push := func(n int) {
		if top >= MAX-1 {
			fmt.Printf("\nstack overflow\n")
		}
		top++
		stack[top] = n
	}

	pop := func() {
		if top < 0 {
			fmt.Printf("\nstack underflow\n")
		}
		top--
	}

	print_stack := func() {
		fmt.Printf("\nstack contents:top->bottom\n")
		for i := top; i >= 0; i-- {
			fmt.Printf("%d  ", stack[i])
		}
	}

	init_stack()
	push(5)
	push(4)
	push(6)
	push(8)
	push(7)
	print_stack()
	pop()
	pop()
	push(1)
	print_stack()
}
