package main

import (
	"fmt"
)

func main() {
	a := make([]int, 2, 4) //공간을 4개로 만들어놓고 2개만 초기화
	//a := []int{1, 2} 공간을 2개 만들고 2개를 초기화
	a[0] = 1
	a[1] = 2

	b := append(a, 3) //a에 공간이 남아서 b가 a와 같은 메모리
	//a에 공간이 남지 않으면 b는 독립된 메모리

	fmt.Printf("%p %p\n", a, b)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 4
	b[1] = 5
	b[2] = 7

	fmt.Println(a)
	fmt.Println(b)

	//a에 공간이 남아서 b의 값을 바꿨는데도 a값이 바뀜
	//a에 공간이 남지 않는다면 b의 값을 바꾸면 a값이 바뀌지 않음
	//a를 b에 복사붙여넣기 하면 아예 다른메모리로 됨 (매우 안전함)
	//b의 값을 바꾸려고 하였는데 a의 값도 바뀌는 버그가 빈번케 발생함
}
