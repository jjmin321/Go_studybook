package main

import "fmt"

var MAX int = 10
var stack = [MAX]int
var top int

func init_stack() {
	top = 0
}

func push(n int) {
	if top >= MAX-1 {
		fmt.Println("stack overflow")
	}
	stack[top] = n
	top++
}

func pop() {
	if top < 0 {
		fmt.Println("stack underflow")
	}
	stack[top] = 0
	top--
}

func print_stack() {
	fmt.Printf("\nstack contents:top->bottom\n")
	for i := top; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}

func main() {
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
