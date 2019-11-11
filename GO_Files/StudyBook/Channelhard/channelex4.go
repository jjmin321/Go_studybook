//20-4

package main

import (
	"fmt"
	"time"
)

func main() {
	//채널(channel) select 구문
	//채널 Select 구문 -> 채널에 값이 수신되면 해당 case문 실행
	//일회성 구문이므로 For(반복문) 안에서 수행
	//default 구문 처리 수행

	//ex1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1: //case <-ch1 과 같지만 <-ch1의 값을 사용하기 위해 num := 을 추가적으로 사용
				fmt.Println("ch1 : ", num)
				// case <-ch1 :
				//     fmt.Println("ch1 : ", <-ch1)
			case str := <-ch2: //case <-ch1 과 같지만 <-ch1의 값을 사용하기 위해 num := 을 추가적으로 사용
				fmt.Println("ch2 : ", str)
				// case <-ch2 :
				//     fmt.Println("ch2 : ", <-ch2)
				// default 사용 안 함
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
