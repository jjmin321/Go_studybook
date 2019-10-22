package main

import "fmt"

func main() {
	//제어문(조건문)
	//IF 문 : Only Boolean
	//Boolean  -> 1, 0 can't use
	//소괄호 사용 X

	var a int = 20

	//ex 1
	if a >= 15 {
		fmt.Println("15 or more")
	}

	/*
		Error 1
		if b >= 25
		{
			fmt.Println("25or more")
		}
	*/

	/*
		Error 2
		if b >= 25
			fmt.Println("25or more")
	*/

	/*
		Error 3
		if c := 1; //true로 바꿔야함 c {
			fmt.Println("True")
		}
	*/

	//ex 2
	if c := 40; c >= 35 {
		fmt.Println("35 or more")
	}

}
