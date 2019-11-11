//20-5

package main

import (
	"fmt"
	"time"
)

func main() {
	//채널(Channel)

	//ex1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1 //값 꺼내기(수신)
			fmt.Println("ch1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!" //값 싣기(발신)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 777:
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
