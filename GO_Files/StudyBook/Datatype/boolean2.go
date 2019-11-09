//1-12

package main

import "fmt"

func main() {
	//ex1
	fmt.Println("ex1 : ", true && true)
	fmt.Println("ex2 : ", true && false)
	fmt.Println("ex3 : ", true || false)
	fmt.Println("ex4 : ", false || false)
	fmt.Println("ex5 : ", !true)
	fmt.Println("ex6 : ", !false)

	//ex2
	num1 := 15
	num2 := 37

	fmt.Println("ex1 : ", num1 < num2)
	fmt.Println("ex2: ", num1 > num2)
	fmt.Println("ex3: ", num1 >= num2)
	fmt.Println("ex4: ", num1 <= num2)
	fmt.Println("ex5: ", num1 == num2)
	fmt.Println("ex6: ", num1 != num2)
}
