package internal

import (
	"fmt"
	"module/internal/domain"
)

type Warrior struct {
	name         string
	passiveBuffs domain.IBuff
	attackSkills map[string]domain.AttackStrategy
	hp           float64
}

var warrior *Warrior

func GetWarrior() *Warrior {
	if warrior == nil {
		fmt.Println("I am Warrior!")
		return &Warrior{
			passiveBuffs: &domain.Buff{},
			attackSkills: make(map[string]domain.AttackStrategy),
			hp:           3000,
		}
	}

	return warrior

}

func (w Warrior) SetAttack(attackName string, strategy domain.AttackStrategy) {
	w.attackSkills[attackName] = strategy
}

func (w Warrior) Damage(attackStrategy string) {
	attackSkill, ok := w.attackSkills[attackStrategy]
	if !ok {
		fmt.Println("No such skill")
		return
	}
	attackSkill.Attack()
}

func (w *Warrior) GetBuff(buff domain.IBuff) {
	w.passiveBuffs = buff
}

func (w *Warrior) GetName() string {
	return w.name
}

func (w Warrior) GetHP() float64 {
	return w.hp
}

func (w Warrior) CheckBuffs() {
	fmt.Println(w.passiveBuffs.GetDescription())
}

func (w *Warrior) update(hp float64) {
	w.hp = w.hp + hp
	fmt.Println("HP increased to: ", w.hp)
}
