//19-1

package main

import (
	"fmt"
	"time"
)

func work1(a chan int) {
	fmt.Println("Work1 : Start ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : ENd ---> ", time.Now())
	a <- 1
}

func work2(a chan int) {
	fmt.Println("Work2 : Start ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : ENd ---> ", time.Now())
	a <- 2
}

func main() {
	//채널(Channel)
	//고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용 : 채널(동기식, 참조 타입)
	//실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	//데이터 전달 자료형 선언 후 사용(지정된 타입만 주고받을 수 있음)
	//interface{} 전달을 통해서 자료형 상관없이 전송 및 수신가능
	//값을 전달(복사 후 : bool, int등), 포인터(슬라이스, 맵) 등을 전달시에는 주의 -> 동기화 사용(Mutex)
	//멀티프로세싱 처리에서 교착상태(경쟁) 주의!
	// <- , -> ( 채널 <- 데이터 : 송신) , ( <- 채널 : 수신 )

	//ex1
	fmt.Println("Main : Start ---> ", time.Now())

	a := make(chan int) //int형 채널 선언
	go work1(a)
	go work2(a)
	<-a
	<-a
	fmt.Println("Main : End ---> ", time.Now())

}
