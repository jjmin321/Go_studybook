package main

import "fmt"

func main() {

	//ex1
	a := 30 / 15
	switch a {
	case 2, 4, 6: //if i is 2, 4, 6
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: //if i is 1, 3, 5
		fmt.Println("a -> ", a, "는 홀수")
	}
	
	//ex2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")

	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")

	case "ruby":
		fmt.Println("ruby")

	case "C":
		fmt.Println("c")
	}
}
