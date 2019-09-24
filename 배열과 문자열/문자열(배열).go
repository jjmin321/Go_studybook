package main

import "fmt"

func main() { //문자열은 배열이고 배열은 메모리다
	//byte 문자열
	s := "Hello 월드" //영어는 1byte 한글은 3byte

	fmt.Println("len(s) = ", len(s))
	for i := 0; i < len(s); i++ { //for문은 1byte씩 돈다
		fmt.Print(string(s[i]), ", ")

	}
}
