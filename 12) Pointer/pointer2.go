package main

import "fmt"

func main() {
	//ex1
	i := 7
	p := &i

	fmt.Println(i, *p)
	fmt.Println(&i, p)
}
