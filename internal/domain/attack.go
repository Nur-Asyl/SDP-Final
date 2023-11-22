package domain

import "fmt"

type AttackStrategy interface {
	Attack()
}

type HookAttack struct{}

func newHookAttack() *HookAttack {
	return &HookAttack{}
}

func (h HookAttack) Attack() {
	fmt.Println("Hook, BAAAAM!")
}

type RootAttack struct{}

func newRootAttack() *RootAttack {
	return &RootAttack{}
}

func (r RootAttack) Attack() {
	fmt.Println("Its smell and melting...")
}

type UltimateDismemberAttack struct{}

func newUltimateDismemberAttack() *UltimateDismemberAttack {
	return &UltimateDismemberAttack{}
}

func (u UltimateDismemberAttack) Attack() {
	fmt.Println("FRESH MEAT!!!")
}
