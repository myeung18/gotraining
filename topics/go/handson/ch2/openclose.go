package ch2

import "fmt"

type Sword struct {
	name string
}

func (s Sword) Damage() int {
	return 2
}

func (s Sword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage to opponents", s.name, s.Damage())
}

type EnchantedSword struct {
	Sword
}

func (e EnchantedSword) Damage() int {
	return 42
}

func (e EnchantedSword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage to opponents", e.name, e.Damage())
}

type Monster struct{}

func Attack(m *Monster, s *Sword) {
}

func testCall() {
	Attack(&Monster{}, &Sword{})
}
