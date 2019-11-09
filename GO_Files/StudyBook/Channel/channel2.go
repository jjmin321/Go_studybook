//19-2

package main

import (
	"fmt"
)

func rangeSum(a int, c chan int) {
	sum := 0

	for i := 1; i <= a; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700, c)
	go rangeSum(500, c)

	//순서대로 데이터 수신(동기) : 채널에서 값 수신 완료될 때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
