//27-1

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

//에러 체크 방식1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

//에러 체크 방식2
func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	//파일 읽기
	//csv 파일 읽기 예제
	//패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	//패키지 Github 주소 : https://github.com/tealeg/xlsx
	//bufio : 파일이 용량이 클 경우 버퍼 사용 권장 (크지 않더라도 그냥 권장)

	//파일 생성
	file, err := os.Open("/users/jejeongmin/documents/go/src/go_studybook/GO_Files/StudyBook/FileRead/sample.csv")
	//리소스 해제
	defer file.Close()

	//CSV Reader 생성
	//rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) //권장

	//CSV 내용 읽기(1) 전혀 쓰지는 않지만 알고 있어도 됨

	row1, err1 := rr.Read()
	errCheck1(err1)
	fmt.Println("CSV Read Example")
	fmt.Println(row1[0], row1[1:7])
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")

	//CSV 내용 읽기(2) 이걸 거의 99프로 사용함.
	rows, err := rr.ReadAll() //전체 row읽기
	errCheck1(err)
	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows[5][1], " : ", rows[2][1], " : ", rows[6][1])
	fmt.Println("≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠≠")

	//Row 단위로 CSV 파일 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s      ", rows[i][j])
		}
		fmt.Println()
	}

}
