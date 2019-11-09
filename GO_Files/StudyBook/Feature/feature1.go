//6-1

package main

import "fmt"

func main() {
	//GO : Ban ambiguous and iffy grammer
	//cant use ++i , only i++

	//ex1
	sum, i := 0, 0

	for i <= 100 {
		//can't use sum += i++
		sum += i
		i++
	}
	fmt.Println("ex1: ", sum)
}
