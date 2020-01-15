package main

import "fmt"

func main() {
	var s student
	s.name = "sh"
	s.class = 1

	s.Score.name = "math"
	s.Score.grade = "A"

	s.ViewScore()

	s.InputScore("과학", "B")
	s.ViewScore()

}

type student struct {
	name  string
	class int

	Score score
}
type score struct {
	name  string
	grade string
}

func (s student) ViewScore() {
	fmt.Println(s.Score)
}

// 값을 변경하는 것이기 때문에 pointer로 호출 해야한다
func (s *student) InputScore(name string, grade string) {
	s.Score.name = name
	//s->Score.name
	s.Score.grade = grade
} //함수는 값이 복사 될 뿐 -> 그렇기 때문에 이미 선언 된 구조체의 값이 변하지는 않음
//해결 방법은 포인터! student 구조체를 받아올 때 포인터로 받아오면 s에 원본 s의 주소값을 가져옴
//즉, s의 값을 변경 할 수 있음


