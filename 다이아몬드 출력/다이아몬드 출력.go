package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 4; i > 0; i-- {
		for j := 4; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < i*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j < i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 5; i >= 1; i-- {
		for j := 5; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 1; j < i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}*/
