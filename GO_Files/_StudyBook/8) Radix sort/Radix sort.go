package main //main 패키지에 포함된다.

import "fmt" //기본적인 출력과 입력 포맷 관련한 기능 제공

func main() { //main 함수
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 0, 5} //0~9 범위의 배열 선언
	fmt.Printf("arr배열의 값은:%d입니다\n", arr)
	temp := [10]int{} //(0에서 9까지 셀 수 있는)최소 10개 크기의 배열을 선언하고 모두 0으로 초기화

	for i := 0; i < len(arr); i++ { //arr배열 분석
		idx := arr[i] //arr 배열 i번째 값을 idx에 넣음
		temp[idx]++
		//temp[idx]에 1을 더함
		//ex)temp[0(arr의 첫 번째 값)]에 1을 더함
		//ex)temp[5(arr의 두 번째 값)]에 1을 더함
		//ex)temp[4(arr의 세 번째 값)]에 1을 더함
		//최종적으로 temp := [10]int{2, 1, 1, 1, 1, 2, 1, 0, 1, 1} 가 됨
	}
	fmt.Printf("temp배열의 값은:%d입니다\n", temp)

	idx := 0                         //변수 idx를 0으로 선언
	for i := 0; i < len(temp); i++ { //temp배열크기만큼 반복
		for j := 0; j < temp[i]; j++ { //temp배열 분석
			arr[idx] = i //temp[0]의 값만큼 0을 arr배열에 입력 (temp[0]는 arr배열에 있는 0의 개수)
			idx++        //arr[1]에 다음 값을 넣을 준비를 함
			//만약 temp[0]의 값이 2라면 arr[0], arr[1]에 모두 0이 입력됨
		}
	}
	fmt.Printf("바뀐 arr배열의 값은:%d입니다\n", arr) //arr배열 출력
}
