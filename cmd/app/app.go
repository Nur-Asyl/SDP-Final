package app

import (
	"fmt"
	"module/internal"
)

func Start() {

	fmt.Println("CREATING HERO:")
	black := internal.GetWarrior()
	white := internal.GetMage()

	black.SetAttack("Hook", internal.HookAttack)
	black.SetAttack("Root", internal.RootAttack)
	black.SetAttack("Ultimate", internal.UltimateDismemberAttack)

	white.SetAttack("Meteor", internal.MeteorCast)
	white.SetAttack("Beem Laser", internal.BeemLaserCast)

	fmt.Println("\n----Warrior BUFFS----")
	black.GetBuff(internal.StrengthAndSpeedBuff)
	black.CheckBuffs()

	fmt.Println("\n----Mage BUFFS----")
	white.GetBuff(internal.SpeedBuff.Buff)
	white.CheckBuffs()

	fmt.Println("\n----Warrior ATTACK----")
	fmt.Println("Hook SKILL:")

	black.Damage("Hook")

	fmt.Println("Root SKILL:")

	black.Damage("Root")

	fmt.Println("ULTIMATE SKILL:")

	black.Damage("Ultimate")

	fmt.Println("\n----MAGE ATTACK----")

	fmt.Println("Meteor Spell:")

	white.Cast("Meteor")

	fmt.Println("Beem Laser Spell:")

	white.Cast("Beem Laser")

	fmt.Println("\n----HP Shrine Black----")

	hpShrine := internal.HPShrine{
		BarManager: &internal.HPShrineManager{},
	}

	fmt.Println("HP: ", black.GetHP())

	hpShrine.BarManager.SetState(100)
	hpShrine.Register(black)
	hpShrine.NotifyAll()

	fmt.Println("HP: ", black.GetHP())

}
