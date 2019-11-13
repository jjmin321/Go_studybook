//21-4

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//뮤텍스(mutex) : 상호 배제 -> Goroutine들이 서로 running time에 서로 영향을 주지 않게, 단독으로 실행되게 하는 기술
	//뮤텍스(mutex) : 여러 고루틴에서 작업하는 공유 데이터 보호
	//RWMutex : 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 X, 읽기, 쓰기 Lock 전부 방어(방지)
	//RMutex : 읽기 Lock -> 읽기 시도 중에 값이 변경 방지, 쓰기 Lock 방어(방지)

	//동기화 사용한 경우 예제
	//쓰기 읽기 동작 순서가 일정하므로 잘못된 오류를 반환할 가능성 없음
	//시스템 전체 cpu 사용

	//ex1
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			//쓰기 뮤텍스 잠금
			mutex.Lock()
			data++
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			//쓰기 뮤텍스 잠금 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i <= 10; i++ {
			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
			//읽기 뮤텍스 잠금 해제
			mutex.RUnlock()

		}
	}()

	go func() {
		for i := 0; i <= 10; i++ {
			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			//읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
