package main

type Player struct {
}

type Monster struct {
}

// 객체에 의존하면 확장 시 어려움을 겪음, attack 가능한 box 라는 객체가 추가 되면 함수가 1+1+3 개가 새로 쓰여져야함
func (p *Player) AttackP(pAttacked *Player) {

}
func (p *Player) AttackM(mAttacked *Monster) {

}
func (m *Monster) AttackP(pAttacked *Player) {

}
func (m *Monster) AttackM(mAttacked *Monster) {

}

// 이런 식으로 interface에 묶이면 box 라는 객체가 새로 생겨도 box에 관한 함수만 추가해주면 됨
type DoAttack interface {
	Attack(BeAttacked)
}
type BeAttacked interface {
	BeAttacked()
}

// 또는 함수에서 애초에 아래와 같이 선언
func Attack(attacker DoAttack, defender BeAttacked) {
	attacker.Attack(defender)
}

// interface 로 묶인 경우 이처럼 간단하게 코드 작성 가능
func (p *Player) Attack(target BeAttacked) {

}
func (m *Monster) Attack(target BeAttacked) {

}

// 새로운 객체가 추가되어도 이렇게 간단히 추가 가능
type Box struct {
}

func (b *Box) Attack(target BeAttacked) {

}
