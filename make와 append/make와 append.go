package main

import (
	"fmt"
)

func main() {
	a := make([]int, 6)

	for i := 0; i < 6; i++ {
		a[i] = i + 1
	}
	a = append(a, 7, 8, 9)

	fmt.Println(a, len(a), cap(a))

}

/*
	길이 6인 배열 a를 만듦
	1, 2, 3, 4, 5, 6 넣음
	배열 a에 7,8,9 추가
	len은 쓰고 있는 배열 길이라서 6에서 3개가 늘어난 9가 되었다.

	cap은 배열 길이가 초과될 경우 2배로 늘리기 때문에 6에서 12로
	최대 배열 길이가 늘어남
*/
