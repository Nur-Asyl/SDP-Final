package domain

type CastAdapter struct {
	m ISpell
}

func (a CastAdapter) Attack() {
	a.m.Cast()
}
