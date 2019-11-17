//28-2

package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 활용
	//ioutil, bufio 등은 io.reader, io.writer 인터페이스를 구현함 -> 즉, Write, Read 메소드를 사용가능
	/*
		type Reader interface{
			Read(p []byte) (n int, err error)
		}
		type Writer interface{
			write(p []byte) (n int, err error)
		}
	*/
	//즉, bufio의 NewReader, NewWriter을 통해서 객체 반환 후 메소드 사용

	//bufio(Buffered io) 패키지
	//https://golang.org/pkg/bufio

	//파일 열기
	//두 번째 매개변수 확인
	//https://golang.org/pkg/os/#pkg-constants

	/* 버퍼
	상태
	a ---------> a ----------> NIL
	b ---------> a, b ----------> NIL
	c ---------> a, b, c ----------> NIL
	d ---------> a, b, c, d ----------> NIL
	e ---------> e ----------> a, b, c, d (a, b, c, d 는 버퍼가 가득 차자 전송됨)
	f ---------> e, f ----------> a, b, c, d
	g ---------> e, f, g ----------> a, b, c, d
	h ---------> e, f, g, h ----------> a, b, c, d
	i ---------> i ----------> a, b, c, d, e, f, g, h
	*/

	file, err := os.OpenFile("users/jejeongmin/documents/go/src/go_studybook/go_files/studybook/fileioutil/test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	// os.O_CREATE|os.O_RDWR -> 없으면 만들고, 있으면 퍼미션 7 제공

	//bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) //Writer 반환(버퍼 사용)
	wt.WriteString("Hello golang!\n File Write Test!\n")
	wt.WriteString("What's your name!\n")

	//에러 체크
	errCheck(err)

	//버퍼 정보 출력
	fmt.Printf("사용한 Buffer Size (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d bytes)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d bytes)\n", wt.Size())

	//버퍼 비우고 반영(버퍼의 내용을 디스크에 반영)
	wt.Flush()
	fmt.Println("쓰기 작업 완료\n")
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")

	rt := bufio.NewReader(file) //Reader 반환
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b) //읽기(ReadLine, ReadByte, ReadBytes 등)
	//rt.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")
	fmt.Println(string(b))
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")
	defer file.Close()

}
