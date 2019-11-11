//20-2

package main

import "fmt"

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i < cnt+1; i++ {
			sum += i
		}
		tot <- sum //tot가 sum의 값을 수신하고 다시 발신을 해야하는데 goroutine을 사용하지 않으면 발신할 곳이 당장 없어서 데드락 상태가 일어남.
	}()

	return tot
}

func main() {
	//채널(channel)
	//채널 또한 함수의 반환 값으로 사용 가능

	//ex1
	c := sum(100) //c := <-tot

	fmt.Println("ex1 : ", <-c)

}
