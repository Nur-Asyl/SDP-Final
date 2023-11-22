package internal

import (
	"fmt"
	"module/internal/domain"
)

type Mage struct {
	name         string
	passiveBuffs domain.IBuff
	attackSkills map[string]domain.AttackStrategy
	hp           float64
}

var mage *Mage

func GetMage() *Mage {

	if mage == nil {
		fmt.Println("I am Mage!")
		return &Mage{
			passiveBuffs: &domain.Buff{},
			attackSkills: make(map[string]domain.AttackStrategy),
			hp:           2000,
		}
	}

	return mage

}

func (m Mage) SetAttack(attackName string, strategy domain.AttackStrategy) {
	m.attackSkills[attackName] = strategy
}

func (m Mage) Cast(attackStrategy string) {
	attackSkill, ok := m.attackSkills[attackStrategy]
	if !ok {
		fmt.Println("No such skill")
		return
	}
	attackSkill.Attack()
}

func (m *Mage) GetBuff(buff domain.IBuff) {
	m.passiveBuffs = buff
}

func (m Mage) CheckBuffs() {
	fmt.Println(m.passiveBuffs.GetDescription())
}

func (m *Mage) GetName() string {
	return m.name
}

func (m Mage) GetHP() float64 {
	return m.hp
}

func (m *Mage) update(hp float64) {
	m.hp = m.hp + hp
	fmt.Println("HP increased to: ", m.hp)
}
