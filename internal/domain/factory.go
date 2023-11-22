package domain

import "fmt"

func GetSpell(spellType string) AttackStrategy {
	if spellType == "meteor" {
		a := CastAdapter{
			m: newMeteorCast(),
		}
		return a
	}
	if spellType == "beemlaser" {
		a := CastAdapter{
			m: newBeemLaserCast(),
		}
		return a
	}
	fmt.Println("No such attack")
	return nil
}

func GetAttack(attackType string) AttackStrategy {
	if attackType == "hook" {
		return newHookAttack()
	}
	if attackType == "root" {
		return newRootAttack()
	}
	if attackType == "ultimateDismember" {
		return newUltimateDismemberAttack()
	}
	fmt.Println("No such attack")
	return nil
}
