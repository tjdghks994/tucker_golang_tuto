package main

import "fmt"

type Bread struct {
	val string
}

type Jam struct {
}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

// String Println 에서 출력을 할 때 struct에 String() 함수가 있으면 그 함수의 값을 출력함
func (b *Bread) String() string {
	return b.val + " : print string func "
}

func (j *Jam) GetVal() string {
	return " + jam"
}

func main() {
	/* OOP 에서는 각 객체의 기능이 중요함
	각 객체의 상태와 기능 그리고 객체 간의 관계가 중요함
	객체는 내부 method와 외부 method를 갖고 있음
	인터페이스(interface): 객체간의 상호작용을 정의한 것 -> 외부 method를 따로 정리 한 것 
	즉, 객체는 상태 + 기능으로 이루어져 있는데, 기능 부분을 따로 떼어내어(종속성을 분리하였다) 표현한 것 = interface 
	즉, 관계만 유지 된다면, 객체는 어느것으로 바뀌어도 괜찮다
	*/ 
	bread := &Bread{val: "bread"}
	jam := &Jam{}

	bread.PutJam(jam)

	fmt.Println(bread)
}
