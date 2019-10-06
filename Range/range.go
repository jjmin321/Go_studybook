package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { // for ~ range 는 pow 크기만큼 for문을 돌린다는 뜻으로 첫 번째 오는 변수에는 배열 i번째, 두 번째 오는 변수에는 i번째의 값을 넣음.
		fmt.Println(i, "=", v)
	}
}
