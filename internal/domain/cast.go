package domain

import "fmt"

type ISpell interface {
	Cast()
}

type MeteorCast struct{}

func newMeteorCast() *MeteorCast {
	return &MeteorCast{}
}

func (m MeteorCast) Cast() {
	fmt.Println("Meteor!, BAAAAM!")
}

type BeemLaserCast struct{}

func newBeemLaserCast() *BeemLaserCast {
	return &BeemLaserCast{}
}

func (m BeemLaserCast) Cast() {
	fmt.Println("pshhhhiiiiiiiiiiiiiiiiiik, PSHIUYYY")
}
