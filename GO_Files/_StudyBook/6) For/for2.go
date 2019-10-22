package main

import "fmt"

func main() {

	//ex1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("ex1: ", sum1)

	//ex2
	sum2, i := 0, 0

	for i <= 100 { //same as while
		sum2 += i
		i++
	}
	fmt.Println("ex2: ", sum2)

	//ex3
	sum3, j := 0, 0

	for { //same as while
		if j > 100 {
			break
		}
		sum3 += j
		j++
	}
	fmt.Println("ex3: ", sum3)

	//ex4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4: ", i, j)
	}

}
