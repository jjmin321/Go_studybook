//20-3

package main

import "fmt"

func receive(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 7777
		close(tot)
	}()
	return tot
}

func total(ch <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range ch {
			a = a + i
		}
		tot <- a
	}()
	return tot
}

func main() {
	//채널(channel)

	//ex1
	c := receive(100)  //receiveOnly함수에서 채널 생성후 c에 꺼내기 전용 채널 tot의 참조값을 꺼내서 집어넣음
	output := total(c) //꺼내기 전용 채널 c를 total함수에 보낸 후
	//fmt.Println(output) //채널은 레퍼런스 타입이기 때문에 값이 출력되지 않음
	fmt.Println(<-output) //값을 꺼내서 출력해야함.
}
