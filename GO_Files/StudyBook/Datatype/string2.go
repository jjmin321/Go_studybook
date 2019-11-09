//1-2

package main

import (
	"fmt"
)

func main() {
	//문자열 표현
	//문자열 : UTF - 8 .

	//ex1
	var str1 string = "Golang"
	var str2 string = "World!"
	var str3 string = "고프로그래밍"

	fmt.Println(str1[0])
	fmt.Println(str1[1])
	fmt.Println(str1[2])
	fmt.Println(str1[3])
	fmt.Println(str1[4])
	fmt.Println(str1[5])

	fmt.Println(str2)
	fmt.Println(str3[0])
	fmt.Println(str3[1])
	fmt.Println(str3[2])
	fmt.Println(str3[3])
	fmt.Println(str3[4])
	fmt.Println(str3[5])

	//ex2
	fmt.Printf("ex 1: %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])

	//ex3
	for i, j := range str1 {
		fmt.Printf("ex3 : %c(%d)\t", j, i)
	}
	fmt.Println()

	for i, j := range str2 {
		fmt.Printf("ex4 : %c(%d)\t", j, i)
	}
}
