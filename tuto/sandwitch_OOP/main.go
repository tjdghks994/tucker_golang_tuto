package main

import "fmt"

// Bread define
type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) { // Jam이라는 관계를 받아오면 됨
	spoon := jam.GetOneSpoon() // Bread와 jam의 관계를 맺어주고 있음
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}

// StrawberryJam define
type StrawberryJam struct {
}
type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberryJam"
}
func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

// OrangeJam define
type OrangeJam struct {
}
type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return " + OrangeJam"
}
func (o *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

// AppleJam define
type AppleJam struct {
}
type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return " + AppleJam"
}
func (o *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

// interface define
type Jam interface {
	GetOneSpoon() SpoonOfJam // 함수명이 GetOneSpoon이며 인풋은 ()이며 return값이 SpoonOfJam 인 함수를 갖고 있는 애들을 묶어줌
}
type SpoonOfJam interface {
	String() string // 함수명이 String이며 인풋은 ()이며 return값이 string 인 함수를 갖는 애들을 묶어줌
}

func main() {
	bread := &Bread{}

	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	// jam 종류를 추가하는데 어려움이 없다, 빵과의 관계로 정의 되어 있기 때문에 빵의 내용을 수정 할 필요 없이 관계(interface)를 유지만 해주면 된다!

	bread.PutJam(jam)

	fmt.Println(bread)
}
