package internal

import "module/internal/domain"

type Hero interface {
	SetAttack(string, domain.AttackStrategy)
	GetBuff(domain.IBuff)
	GetName() string
	CheckBuffs()
	update(float64)
}
