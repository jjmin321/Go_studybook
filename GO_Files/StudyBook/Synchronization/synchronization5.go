//21-5

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//고루틴 동기화 객체
	//동기화 상태(조건) 메소드 사용
	//Wait, notify , notifyAll : 기타 언어
	//Wait, Signal , Broadcast

	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	c := make(chan int, 5) //비동기 버퍼 채널 선언

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Wating : ", n)
			condition.Wait() //고루틴 대기(멈춤)
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("received : ", <-c)
	}

	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Wake Goroutine(Signal) : ", i)
	// 	condition.Signal()
	// 	mutex.Unlock()
	// }

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(1 * time.Second)
}
