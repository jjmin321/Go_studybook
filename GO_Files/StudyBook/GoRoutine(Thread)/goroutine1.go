//18-1

package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 function start!", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 function end", time.Now())
}

func exe2() {
	fmt.Println("exe2 function start!", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 function end", time.Now())
}

func exe3() {
	fmt.Println("exe3 function start!", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 function end", time.Now())
}

func main() {
	//고루틴(goroutine)
	//타 언어의 쓰레드(Thread)와 비슷한 기능
	//생성 방법 매우 간단, 리소스 매우 적게 사용 가능
	//즉, 수많은 고루틴 동시 생성 실행 가능
	//비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통한 통신 가능
	//공유메모리 사용 시에 정확한 동기와 코딩 필요
	//싱글루틴에 비해 항상 빠른 처리 결과는 아니다.

	//멀티 고루틴(쓰레드) 장점과 단점
	//장점 : 응답성 향상, 자원공유를 효율적으로 활용 사용, 작업이 분리되어 코드 간결
	//단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스의 사이드이펙트 , 성능 저하, 동기화 코딩 반드시 숙지
	//      데드락(교착 상태)...

	exe1() //가장 먼저 실행(일반적인 실행 흐름)

	fmt.Println("Main Routine start!", time.Now())
	go exe2()
	go exe3()
	time.Sleep(1 * time.Second)
	fmt.Println("Main Routine End", time.Now())
}
