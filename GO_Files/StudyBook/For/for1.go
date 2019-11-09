//3-1

package main

import "fmt"

var loc = []string{"Seoul", "Busan", "Incheon"} //ex3 Global variable

func main() {
	//Loop - for
	//Go only offer for
	//study various usage

	//ex1
	for i := 0; i < 5; i++ {
		fmt.Println("ex 1 : ", i)
	}

	//error1
	/* for i := 0; i < 5; i++
	 {
		 fmt.Println("ex1 : ", i)
	 }
	*/
	//error2
	/* for i := 0; i < ; i++
		 fmt.Println("ex1 : ", i)
	}
	*/

	//ex2 infinite loop
	for {
		fmt.Println("ex2 : Infinite loop!")
	}

	//ex3 Range
	//loc := []string{"Seoul", "Busan", "Incheon"}	//local variable
	for index, value := range loc {
		fmt.Println("ex3 : ", index, value)

	}
}
