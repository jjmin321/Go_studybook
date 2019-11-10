//19-4

package main

import (
	"fmt"
	"runtime"
)

func main() {
	//채널(channel)

	//ex1(비동기: 버퍼 사용)
	//버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동, 수신 -> 비어있으면 대기
	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 8) //버퍼가 2개인 boolean형 channel, 즉 용량이 2개인 것, 받을 때 2개 있어야 보냄.
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true //가득 차면 대기
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch //비어있으면 대기
		fmt.Println("Main : ", i)
	}
}
