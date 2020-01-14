package main

import "fmt"

type Bread struct {
	val string
}

type strawberryJam struct {
	opened bool
}

type SpoonOfJam struct {
}

type Sandwitch struct {
	val string
}

func GetsBreads(v int) []*Bread {
	breads := make([]*Bread, v)

	for i := 0; i < v; i++ {
		breads[i] = &Bread{val: "bread"}
	}

	return breads
}

func OpenStrawberryJam(jam *strawberryJam) {
	jam.opened = true
}

func GetOneSpoon(_ *strawberryJam) *SpoonOfJam {
	return &SpoonOfJam{}
}

func PutJamOnBread(bread *Bread, spoon *SpoonOfJam) {
	bread.val += " + strawberry Jam"
}

func MakeSandwitch(bread []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(bread); i++ {
		sandwitch.val += bread[i].val + " + "
	}

	return sandwitch
}

func main() {
	// 절차적 프로그래밍: 과정이 눈에 쉽게 보인다!
	// 단점: 유지보수가 아주 힘들다 ex 딸기잼 -> 사과잼 으로 바꾸려면 기존 함수를 사용할 수 없기 때문에 모든 함수를 다 바꿔야함
	
	// 1. 빵 두개 꺼낸다
	breads := GetsBreads(2)

	jam := &strawberryJam{}

	// 2. 딸기잼 뚜껑을 연다
	OpenStrawberryJam(jam)

	// 3. 딸기잼을 한 스푼 뜬다
	spoon := GetOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다
	PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
