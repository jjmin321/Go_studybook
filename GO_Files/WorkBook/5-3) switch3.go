package main

import "fmt"

func main() {

	// 1.switch와 fallthrough를 사용하여 케이스에 해당하는 값이 출력되면 아래 케이스 1개를 더 출력시키는 프로그램을 만드세요.
	switch e := "go"; e {
	
		fmt.Println("Java")


		fmt.Println("go")
	

		fmt.Println("python")

	
		fmt.Println("ruby")

	
		fmt.Println("c")
	}
	// 출력 결과 : go, python
}