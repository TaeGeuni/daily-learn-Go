package main

type Attacker interface {
	Attack()
}

type Test struct{}

func (t *Test) Attack() {
}

func main() {
	var att Attacker // ❶ 기본값은 nil입니다.

	// var t Test
	// att = &t

	att.Attack() // ❷ att가 nil이기 때문에 런타임 에러가 발생합니다.
}
