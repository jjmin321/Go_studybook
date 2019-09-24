package main

import "fmt"

func main() {
	var s []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	fmt.Println(s, len(s), cap(s))

	s = append(s, 400, 500, 600, 700)

	fmt.Println(s, len(s), cap(s))
}
