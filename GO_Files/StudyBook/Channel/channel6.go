//19-6

package main

import (
	"fmt"
)

func main() {
	//채널(Channel)
	//Close : 채널 닫기

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good!"
		}
	}()
	value1, ok1 := <-ch
	fmt.Println("value1 : ", value1, "ok1 : ", ok1)
	value2, ok2 := <-ch
	fmt.Println("value2 : ", value2, "ok2 : ", ok2)
	value3, ok3 := <-ch
	fmt.Println("value3 : ", value3, "ok3 : ", ok3)
	close(ch) //채널 닫기
}
