package main

import "fmt"

func main() {
	//ex1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1: ", i, j)
		}
	}

	//ex2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2: ", i)
	}

	//ex3
Loop2:
	//error(Loop 레이블 밑에 관련이 없는 소스코드가 있을 때)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3: ", i, j)
		}
	}

}
