//20-1

package main

import (
	"fmt"
	"time"
)

func sendOnly(c chan<- int, cnt int) { // 이 함수에서는 int형 데이터를 채널에 보낼 수만 있다.
	for i := 0; i < cnt; i++ {
		c <- i
	}

	c <- 777
	// fmt.Println(<-c) //발신전용 함수(c chan<- int)에서 수신(<-c)시 에러 발생
}

func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}

	fmt.Println(<-c)
}

func main() {
	//채널(channel)
	//함수 등의 매개변수에 수신 및 발신 전용 채널 지정 가능
	//전용 채널 설정 후 방향이 다를 경우 예외 발생
	//발신 전용 channel <- 데이터형
	//수신 전용 <- channel
	//매개 변수를 통해서 전용 채널인지 확인할 수 있다.
	//채널또한 함수의 반환 값으로 사용할 수 있다.

	//ex1
	c := make(chan int)

	go sendOnly(c, 100) //발신 전용
	go receiveOnly(c)   //수신 전용

	time.Sleep(2 * time.Second) //2초

}
