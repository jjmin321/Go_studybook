package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8] //b에 a의 배열 4번째 수부터 7번째 수까지 넣음
	b[0] = 1    //b가 a를 가리키고 있기 때문에 a의 4번째 수도 바뀜
	b[1] = 2
	fmt.Println(a)
	c := a[4:] //c에 a의 배열 4번째 수부터 마지막 수까지 넣음
	d := a[:4] //d에 a의 첫 수부터 배열 4번째 수까지 넣음
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
