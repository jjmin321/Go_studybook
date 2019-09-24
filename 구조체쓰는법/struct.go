package main //main패키지에 포함된다

import "fmt" //기본적인 출력과 입력 포맷 관련한 기능 제공

type Student struct { //Student 타입의 struct 정의
	class  int
	name   string
	number int
	score  Score //Score 타입의 struct를 의미하는 score 선언
}

type Score struct { //Score 타입의 struct 정의
	name  string
	grade string
}

func ViewScore(jungmin Student) { //ViewScore이라는 메소드 정의
	fmt.Println(jungmin.score) //jungmin.ViewScore을 선언했을 때 나의  정보를 출력
}

func main() {
	var jungmin Student
	jungmin.class = 1
	jungmin.name = "제정민"
	jungmin.number = 19

	jungmin.score.name = "수학"
	jungmin.score.grade = "87"

	ViewScore(jungmin)
}
