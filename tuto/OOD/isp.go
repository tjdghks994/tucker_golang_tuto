package main

type Actor struct {
}

func (a *Actor) Move() {

}

func (a *Actor) Attack() {

}
func (a *Actor) Talk() {

}

// ActorInteface 이렇게 한 interface에 여러 기능을 넣지 말아라, 의존성을 발생 시킴
type ActorInteface interface {
	Move()
	Attack()
	Talk()
}

// 이렇게 각각 interface를 구성하는 것이 낫다, "단일 책임 원칙"을 지키기 좋음
type Moveable interface {
	Move()
}
type Attackable interface {
	Attack()
}
type Talkable interface {
	Talk()
}
